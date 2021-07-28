/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    errcode
 * @Date:    2021/5/28 3:35 下午
 * @package: errcode
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package errcode

import (
	"fmt"
	"github.com/jageros/hawos/protos/pb"
)

var (
	InternalErr  = &err{-1, "服务器内部错误"}     // 服务器内部错误
	Success      = &err{200, "successful"} // 成功
	VerifyErr    = &err{401, "验证失败"}       // 验证失败
	InvalidParam = &err{412, "无效参数"}       // 参数无效
)

// IErr 自定义错误接口
type IErr interface {
	Error() string
	ErrMsg() string
	Code() int32
	ECode() pb.ErrCode
}

type err struct {
	code   int32
	errMsg string
}

func (e *err) Error() string {
	return fmt.Sprintf("ErrCode=%d ErrMsg=%s", e.code, e.ErrMsg())
}

func (e *err) Code() int32 {
	return e.code
}

func (e *err) ErrMsg() string {
	if e.errMsg == "" {
		return "Unknown Errcode"
	}
	return e.errMsg
}

func (e *err) ECode() pb.ErrCode {
	return pb.ErrCode(e.Code())
}

// =========

// New 创建一个错误码，业务逻辑上的错误，错误码使用1000-1999
func New(code int32, errMsg string) IErr {
	return &err{
		code:   code,
		errMsg: errMsg,
	}
}

func WithErr(err_ error) IErr {
	if err_ == nil {
		return nil
	}
	return &err{
		code:   -2,
		errMsg: err_.Error(),
	}
}

func WithErrcode(code int32, err_ error) IErr {
	err2 := &err{
		code: code,
	}
	if err_ != nil {
		err2.errMsg = err_.Error()
	}
	return err2
}

//// 服务内部错误码
//var (
//	InternalErr  = &err{-100, "服务器内部错误"}     // 服务器内部错误
//	Success      = &err{-101, "successful"} // 成功
//	VerifyErr    = &err{-102, "验证失败"}       // 验证失败
//	InvalidParam = &err{-103, "无效参数"}       // 参数无效
//)
