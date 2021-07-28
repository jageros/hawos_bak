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
	"github.com/jageros/hawos/protos/pbf"
)

type IQueue interface {
	PushProtoMsg(msgId int32, arg interface{}, target *pbf.Target) error
	Push(msg *pbf.QueueMsg) error
}
