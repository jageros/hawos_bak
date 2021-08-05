/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    pull
 * @Date:    2021/7/5 6:49 下午
 * @package: kafka
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/jageros/hawos/log"
	"github.com/jageros/hawos/protos/pbf"
	"github.com/jageros/hawos/transport"
	"github.com/jageros/hawos/transport/ws"
	"time"
)

var consumerHandler ws.Writer

func SetConsumerHandles(w ws.Writer) {
	consumerHandler = w
}

type Consumer struct {
	*transport.BaseServer
	topic   string
	cg      sarama.ConsumerGroup
	groupId string
}

func NewConsumer(ctx context.Context, topic string, opfs ...transport.SvrOpFn) *Consumer {
	csr := &Consumer{
		BaseServer: transport.NewBaseServer(ctx, opfs...),
		topic:      topic,
	}

	return csr
}

func (c *Consumer) Init(id, name string) error {
	c.Options.ID = id
	c.Options.Name = name
	c.groupId = id
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V0_11_0_2

	cli, err := sarama.NewClient(c.Options.Endpoints, config)

	//csr, err := sarama.NewConsumer(c.options.Endpoints, config)
	if err != nil {
		return err
	}

	cg, err := sarama.NewConsumerGroupFromClient(c.groupId, cli)

	if err != nil {
		return err
	}

	c.cg = cg
	return nil
}

func (c *Consumer) Serve() (err error) {
	for {
		select {
		case <-c.Ctx.Done():
			return c.Ctx.Err()
		default:
			err = c.cg.Consume(c.Ctx, []string{c.topic}, &handler{name: c.groupId})
			if err != nil {
				log.Errorf("Consume err: %v", err)
				return err
			}
		}
	}
}

func (c *Consumer) Stop() {
	if c.cg != nil {
		c.cg.Close()
	}

	c.Cancel()
}

type handler struct {
	name string
}

func (h *handler) Setup(assignment sarama.ConsumerGroupSession) error { return nil }
func (h *handler) Cleanup(assignment sarama.ConsumerGroupSession) error {
	return nil
}
func (h *handler) ConsumeClaim(assignment sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	end := time.Now()
	for msg := range claim.Messages() {
		start := time.Now()
		take2 := start.Sub(end)
		if take2 > time.Second {
			log.Infof("Kafka Recv Msg take: %s", take2.String())
		}
		if msg == nil {
			log.Infof("kafka ConsumeClaim recv msg=nil")
			continue
		}
		kmsg := &pbf.QueueMsg{}
		err := kmsg.Unmarshal(msg.Value)
		if err != nil {
			log.Errorf("kafka Unmarshal msg err=%v", err)
			continue
		}

		target := new(ws.Target).CopyPbTarget(kmsg.Targets)
		err = consumerHandler.Write(kmsg.Data, target)
		if err != nil {
			log.Errorf("kafka ConsumeClaim msg handle return err=%v", err)
		}

		assignment.MarkMessage(msg, "") // 确认消息
		end = time.Now()
		take := end.Sub(start)
		if take > time.Second {
			log.Infof("Kafka Consume Msg take: %s", take.String())
		}
	}
	return nil
}
