/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    conf
 * @Date:    2021/6/18 3:24 下午
 * @package: conf
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package conf

import "time"

const (
	JWT_SECRET = "64981e500279991b18c4ca082fc18d44"

	RPC_CALL_TIMEOUT      = time.Second * 2  // rpc 等待连接准备好超时时间
	HTTP_SHUTDOWN_TIMEOUT = time.Second * 30 // http shutdown 超时时间
	READ_TIMEOUT          = time.Second * 5  //  服务读超时（http/rpc/websocket）
	WRITE_TIMEOUT         = time.Second * 5  // 服务写超时 (http/rpc/websocket）
	CLOSE_TIMEOUT         = time.Second * 30 //  stop context timeout
	TOKEN_TIMEOUT         = time.Hour * 24   // token过期时间
)