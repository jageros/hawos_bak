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
	errcode2 "github.com/jageros/hawos/errcode"
	"github.com/jageros/hawos/log"
	"github.com/jageros/hawos/protos/meta"
	_ "github.com/jageros/hawos/protos/meta"
	"github.com/jageros/hawos/protos/pb"
	transport2 "github.com/jageros/hawos/transport"
	httpc2 "github.com/jageros/hawos/transport/http/httpc"
	queues2 "github.com/jageros/hawos/transport/queues"
	"github.com/nsqio/go-nsq"
	"math/rand"
	"sync"
)

var _ queues2.IQueue = &Producer{}

type Producer struct {
	*transport2.BaseServer
	topic string
	pd    *nsq.Producer
	cfg   *nsq.Config
	clk   *sync.Mutex
}

func (p *Producer) getNodeAddr() (string, error) {
	idx := rand.Intn(len(p.Options.Endpoints))
	url := fmt.Sprintf("http://%s/nodes", p.Options.Endpoints[idx])
	resp, err := httpc2.Request(httpc2.GET, url, httpc2.FORM, nil, nil)
	if err != nil {
		return "", err
	}
	pds := resp["producers"].([]interface{})
	pdn := len(pds)
	if pdn <= 0 {
		return "", errcode2.New(11, "无可用NSQ节点")
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

func (p *Producer) PushProtoMsg(msgId pb.MsgID, arg interface{}, target *pb.Target) error {
	im, err := meta.GetMeta(msgId.ID())
	if err != nil {
		return err
	}
	data, err := im.EncodeArg(arg)
	if err != nil {
		return err
	}
	resp := &pb.Response{
		MsgID:   msgId,
		Code:    pb.ErrCode_Success,
		Payload: data,
	}
	data2, err := resp.Marshal()
	if err != nil {
		return err
	}
	msg := &pb.QueueMsg{
		Data:    data2,
		Targets: target,
	}
	return p.Push(msg)
}

func (p *Producer) Push(msg *pb.QueueMsg) error {
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

func NewProducer(ctx context.Context, topic string, opfs ...transport2.SvrOpFn) (*Producer, error) {
	p := &Producer{
		BaseServer: transport2.NewBaseServer(ctx, opfs...),
		topic:      topic,
		clk:        &sync.Mutex{},
	}

	p.Options.Protocol = transport2.Nsq

	p.cfg = nsq.NewConfig()
	err := p.connectToNsqd()
	return p, err
}

func (p *Producer) Stop() {
	p.Cancel()
	p.pd.Stop()
}
