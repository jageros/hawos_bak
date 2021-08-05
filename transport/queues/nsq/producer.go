/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    publish
 * @Date:    2021/7/2 3:51 下午
 * @package: nsq
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package nsq

import (
	"context"
	"fmt"
	"github.com/jageros/hawos/errcode"
	"github.com/jageros/hawos/log"
	"github.com/jageros/hawos/protos/meta"
	"github.com/jageros/hawos/protos/pbf"
	"github.com/jageros/hawos/transport"
	"github.com/jageros/hawos/transport/http/httpc"
	"github.com/jageros/hawos/transport/queues"
	"github.com/nsqio/go-nsq"
	"math/rand"
	"sync"
	"time"
)

var _ queues.IQueue = &Producer{}

type Producer struct {
	*transport.BaseServer
	topic string
	pd    *nsq.Producer
	cfg   *nsq.Config
	clk   *sync.Mutex
}

func (p *Producer) getNodeAddr() (string, error) {
	idx := rand.Intn(len(p.Options.Endpoints))
	url := fmt.Sprintf("http://%s/nodes", p.Options.Endpoints[idx])
	resp, err := httpc.Request(httpc.GET, url, httpc.FORM, nil, nil)
	if err != nil {
		return "", err
	}
	pds := resp["producers"].([]interface{})
	pdn := len(pds)
	if pdn <= 0 {
		return "", errcode.New(101, "无可用NSQ节点")
	}
	idx = rand.Intn(len(pds))
	pd := pds[idx].(map[string]interface{})
	addr := fmt.Sprintf("%v:%v", pd["broadcast_address"], pd["tcp_port"])
	log.Debugf("NsqAddr=%s", addr)
	return addr, nil
}

func (p *Producer) connectToNsqd() error {
	p.clk.Lock()
	defer p.clk.Unlock()

	if p.pd != nil {
		err := p.pd.Ping()
		if err == nil {
			return nil
		}
		p.pd.Stop()
	}

	addr, err := p.getNodeAddr()
	if err != nil {
		return err
	}
	pd, err := nsq.NewProducer(addr, p.cfg)
	if err != nil {
		return err
	}

	p.pd = pd
	return nil
}

func (p *Producer) PushProtoMsg(msgId int32, arg interface{}, target *pbf.Target) error {
	start := time.Now()
	im, err := meta.GetMeta(msgId)
	if err != nil {
		return err
	}
	data, err := im.EncodeArg(arg)
	if err != nil {
		return err
	}
	resp := &pbf.Response{
		MsgID:   msgId,
		Code:    errcode.Success.Code(),
		Payload: data,
	}
	data2, err := resp.Marshal()
	if err != nil {
		return err
	}
	msg := &pbf.QueueMsg{
		Data:    data2,
		Targets: target,
	}
	err = p.Push(msg)
	take := time.Now().Sub(start).String()
	log.Debugf("Nsq Push Msg take: %s", take)
	return err
}

func (p *Producer) Push(msg *pbf.QueueMsg) error {
	data, err := msg.Marshal()
	if err != nil {
		return err
	}
	err = p.pd.Publish(p.topic, data)
	if err != nil {
		err = p.connectToNsqd()
		if err != nil {
			return err
		}
		err = p.pd.Publish(p.topic, data)
	}
	return err
}

func NewProducer(ctx context.Context, topic string, opfs ...transport.SvrOpFn) (*Producer, error) {
	p := &Producer{
		BaseServer: transport.NewBaseServer(ctx, opfs...),
		topic:      topic,
		clk:        &sync.Mutex{},
	}

	p.Options.Protocol = transport.Nsq

	p.cfg = nsq.NewConfig()
	err := p.connectToNsqd()
	return p, err
}

func (p *Producer) Stop() {
	p.Cancel()
	p.pd.Stop()
}
