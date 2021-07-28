/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    kafka
 * @Date:    2021/7/5 9:52 上午
 * @package: kafka
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/jageros/hawos/internal/pkg/log"
	"github.com/jageros/hawos/protos/meta"
	"github.com/jageros/hawos/protos/pb"
	transport2 "github.com/jageros/hawos/transport"
	queues2 "github.com/jageros/hawos/transport/queues"
)

var _ queues2.IQueue = &Producer{}

type Producer struct {
	*transport2.BaseServer
	topic string
	pd    sarama.AsyncProducer
}

func NewProducer(ctx context.Context, topic string, opfs ...transport2.SvrOpFn) *Producer {
	pd := &Producer{
		BaseServer: transport2.NewBaseServer(ctx, opfs...),
		topic:      topic,
	}

	pd.Options.Protocol = transport2.Kafka

	if len(pd.Options.Endpoints) <= 0 {
		pd.Options.Endpoints = []string{fmt.Sprintf("%s:%d", pd.Options.Ip, pd.Options.Port)}
	}

	return pd
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

	msgData, err := resp.Marshal()
	if err != nil {
		return err
	}

	msg := &pb.QueueMsg{
		Data:    msgData,
		Targets: target,
	}

	return p.Push(msg)
}

func (p *Producer) Push(msg *pb.QueueMsg) error {

	byData, err := msg.Marshal()

	if err != nil {
		return err
	}

	msg_ := &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(msg.Targets.GroupId),
		Value: sarama.ByteEncoder(byData),
	}
	go func() {
		p.pd.Input() <- msg_
	}()
	return nil
}

func (p *Producer) Init(id, name string) error {
	p.Options.ID = id
	p.Options.Name = name
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	producer, err := sarama.NewAsyncProducer(p.Options.Endpoints, config)
	if err != nil {
		return err
	}
	p.pd = producer
	return nil
}

func (p *Producer) Serve() error {
	var offset int64 = -1
	for {
		select {
		case <-p.Ctx.Done():
			return p.Ctx.Err()

		case errMsg := <-p.pd.Errors():
			if offset != errMsg.Msg.Offset {
				p.pd.Input() <- errMsg.Msg
				offset = errMsg.Msg.Offset
			}
			log.Infof("Kafka Error Msg: %v", errMsg.Err)

		case msg := <-p.pd.Successes():
			offset = msg.Offset
			//log.Debugf("kafka successful partition=%d offset=%d", msg.Partition, msg.Offset)
		}
	}
}
func (p *Producer) Stop() {
	p.pd.AsyncClose()
	p.Cancel()
}
