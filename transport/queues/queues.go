/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    queues
 * @Date:    2021/7/21 10:45 上午
 * @package: queues
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package queues

import (
	"context"
	"github.com/jageros/hawos/errcode"
	"github.com/jageros/hawos/protos/pbf"
	"github.com/jageros/hawos/transport"
	"github.com/jageros/hawos/transport/queues/kafka"
	"github.com/jageros/hawos/transport/queues/nsq"
	"github.com/jageros/hawos/transport/ws"
)

type IQueue interface {
	PushProtoMsg(msgId int32, arg interface{}, target *pbf.Target) error
	Push(msg *pbf.QueueMsg) error
	transport.IServer
}

func NewProducer(ctx context.Context, topic string, opfs ...transport.SvrOpFn) (IQueue, error) {
	op := &transport.Option{}
	for _, opf := range opfs {
		opf(op)
	}
	switch op.Protocol {
	case transport.Nsq:
		return nsq.NewProducer(ctx, topic, opfs...)
	case transport.Kafka:
		return kafka.NewProducer(ctx, topic, opfs...), nil
	default:
		return nil, errcode.New(-1, "未知队列类型")
	}
}

func NewConsumer(ctx context.Context, topic string, w ws.Writer, opfs ...transport.SvrOpFn) (transport.IServer, error) {
	op := &transport.Option{}
	for _, opf := range opfs {
		opf(op)
	}
	switch op.Protocol {
	case transport.Nsq:
		csr := nsq.NewConsumer(ctx, topic, opfs...)
		csr.RegistryHandler(w)
		return csr, nil
	case transport.Kafka:
		kafka.SetConsumerHandles(w)
		return kafka.NewConsumer(ctx, topic, opfs...), nil
	default:
		return nil, errcode.New(-1, "未知队列类型")
	}
}
