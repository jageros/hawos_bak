/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    header
 * @Date:    2021/6/23 1:49 下午
 * @package: https
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package httpc

import (
	"github.com/jageros/hawos/consts"
	"net/http"
)

//const (
//	HTTP_HD_APP_TOKEN   = "X-App-Token"
//	HTTP_HD_APP_UID     = "X-App-Uid"
//	HTTP_HD_CLIENT_TYPE = "X-Client-Type"
//	HTTP_HD_APPID       = "X-Appid"
//	HTTP_HD_REQUEST_ID  = "X-App-Request-Id" // 用来进行bug查找和问题跟踪用的
//	HTTP_HD_APP_VERSION = "X-App-Version"
//)

func DefaultHeader() map[string]string {
	return map[string]string{
		consts.HTTP_HD_APPID:       "10004",
		consts.HTTP_HD_CLIENT_TYPE: "1",
		consts.HTTP_HD_APP_VERSION: "0",
		consts.HTTP_HD_REQUEST_ID:  "-1",
	}
}

func SetHeader(req *http.Request, arg map[string]string) {
	header := DefaultHeader()
	for key, val := range arg {
		header[key] = val
	}
	for key, val := range header {
		req.Header.Set(key, val)
	}
}

func GetHeader(arg map[string]string) http.Header {
	h := http.Header{}
	header := DefaultHeader()
	for key, val := range arg {
		header[key] = val
	}
	for key, val := range header {
		h.Set(key, val)
	}
	return h
}
