/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    nsq
 * @Date:    2021/7/2 3:28 下午
 * @package: nsq
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package nsq

import (
	"context"
	"github.com/jageros/hawos/log"
	"github.com/jageros/hawos/protos/pbf"
	"github.com/jageros/hawos/transport"
	"github.com/jageros/hawos/transport/ws"
	"github.com/nsqio/go-nsq"
	"time"
)

type Consumer struct {
	topic   string
	channel string
	csr     *nsq.Consumer
	handler ws.Writer
	*transport.BaseServer
}

func NewConsumer(ctx context.Context, topic string, opts ...transport.SvrOpFn) *Consumer {
	csr := &Consumer{
		topic:      topic,
		BaseServer: transport.NewBaseServer(ctx, opts...),
	}
	return csr
}

func (c *Consumer) RegistryHandler(w ws.Writer) {
	c.handler = w
}

func (c *Consumer) HandleMessage(msg *nsq.Message) error {
	start := time.Now()
	if c.handler == nil {
		log.Errorf("Nsq Consumer Handle = nil")
		return nil
	}

	arg := &pbf.QueueMsg{}
	err := arg.Unmarshal(msg.Body)
	if err != nil {
		log.Errorf("Nsq Consumer Unmarshal err: %v", err)
		return err
	}

	target := new(ws.Target).CopyPbTarget(arg.Targets)
	log.Debugf("Nsq consumer write msg to client, Target=%+v", target)
	err = c.handler.Write(arg.Data, target)
	take := time.Now().Sub(start)
	log.Debugf("Nsq Consumer Msg take: %s", take)
	return err
}

func (c *Consumer) Init(id, name string) (err error) {
	c.Options.ID = id
	c.Options.Name = name
	c.channel = id
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	c.csr, err = nsq.NewConsumer(c.topic, c.channel, cfg)
	if err != nil {
		return
	}
	c.csr.SetLogger(nil, 0)
	c.csr.AddHandler(c)

	err = c.csr.ConnectToNSQLookupds(c.Options.Endpoints)

	return
}
func (c *Consumer) Serve() error {
	select {
	case <-c.Ctx.Done():
		return c.Ctx.Err()
	case i := <-c.csr.StopChan:
		log.Infof("Consumer StopChan=%d", i)
	}
	return nil
}

func (c *Consumer) Stop() {
	c.Cancel()
	if c.csr != nil {
		c.csr.Stop()
	}
}
