/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    consts
 * @Date:    2021/6/17 6:41 下午
 * @package: consts
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package consts

const (
	BaseUid = 81

	QueueTopic = "Hawos"
)

// App Name
const (
	APP_DEMO  = "demo"
	APP_HELLO = "hello"

	APP_CONFIG = "config"
	APP_FRONTEND = "frontend"
	APP_CHAT     = "chat"
)

// http header
const (
	HTTP_HD_APP_TOKEN   = "X-App-Token"
	HTTP_HD_APP_UID     = "X-App-Uid"
	HTTP_HD_CLIENT_TYPE = "X-Client-Type"
	HTTP_HD_APPID       = "X-Appid"
	HTTP_HD_REQUEST_ID  = "X-App-Request-Id" // 用来进行bug查找和问题跟踪用的
	HTTP_HD_APP_VERSION = "X-App-Version"
)