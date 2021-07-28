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
	"github.com/jageros/hawos/protos/pb"
)

type IQueue interface {
	PushProtoMsg(msgId pb.MsgID, arg interface{}, target *pb.Target) error
	Push(msg *pb.QueueMsg) error
}
