/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    utils
 * @Date:    2021/5/28 3:29 下午
 * @package: http
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package http

import (
	"github.com/gin-gonic/gin"
	errcode2 "github.com/jageros/hawos/errcode"
	jwt2 "github.com/jageros/hawos/jwt"
	"github.com/jageros/hawos/log"
	"net/http"
	"net/url"
	"time"
)

func DecodeUrlVal(c *gin.Context, key string) string {
	val, err := url.ParseQuery(key + "=" + c.Query(key))
	if err == nil {
		return val[key][0]
	}
	return ""
}

func PkgMsgWrite(c *gin.Context, data interface{}) {
	code := errcode2.Success
	dataMap := gin.H{"code": code.Code(), "msg": code.ErrMsg()}
	if data != nil {
		dataMap["data"] = data
	}
	log.Debugf("Require successful and write to client msg=%v", dataMap)
	c.JSON(http.StatusOK, dataMap)
}

func ErrInterrupt(c *gin.Context, err errcode2.IErr) {
	log.Errorf("ErrorInterrupt %s", err.Error())
	c.JSON(http.StatusOK, gin.H{"code": err.Code(), "msg": err.ErrMsg()})
	c.Abort()
}

func CheckToken(c *gin.Context) {
	token := DecodeUrlVal(c, "token")
	if token == "" {
		ErrInterrupt(c, errcode2.VerifyErr)
		return
	}
	claims, err := jwt2.ParseToken(token)
	if err != nil {
		log.Infof("ParseToken err: %v", err)
		ErrInterrupt(c, errcode2.VerifyErr)
		return
	}
	if time.Now().Unix() > claims.ExpiresAt {
		ErrInterrupt(c, errcode2.VerifyErr)
		return
	}
	c.Next()
}

func RefreshToken(c *gin.Context) (token string, err error) {
	token = DecodeUrlVal(c, "token")
	if token == "" {
		log.Infof("Get token from url error, not has token value.")
		err = errcode2.VerifyErr
		return
	}
	token, err = jwt2.RefreshToken(token)
	return
}
