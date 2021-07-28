/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    wsapp
 * @Date:    2021/6/10 11:39 上午
 * @package: ws
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package protoc

import (
	errcode2 "github.com/jageros/hawos/errcode"
	"github.com/jageros/hawos/log"
	"github.com/jageros/hawos/protos/meta"
	"github.com/jageros/hawos/protos/pb"
)

var (
	agentRpcHandlers = make(map[int32]AgentRpcHandler)
	msgIds           []int32
)

type AgentRpcHandler func(agent *Agent, arg interface{}) (interface{}, errcode2.IErr)
type AgentRespHandler func(agent *Agent, arg interface{}) errcode2.IErr

func RegisterAgentRpcHandler(msgID meta.IMsgID, handler AgentRpcHandler) {
	agentRpcHandlers[msgID.ID()] = handler
	msgIds = append(msgIds, msgID.ID())
}

func MsgIDs() []int32 {
	return msgIds
}

// OnClientRpcCall req: *pb.Request  resp: *pb.Response
//func OnClientRpcCall(sessID string, data []byte, w io.Writer) {
//	arg := new(pb.Request)
//	err := arg.Unmarshal(data)
//	if err != nil {
//		log.Infof("OnClientRpcCall Unmarshal arg err: %v", err)
//		return
//	}
//	pa := getAgent(sessID)
//	var resp = new(pb.Response)
//	resp.MsgID = arg.MsgID
//	defer func() {
//		respData, err := resp.Marshal()
//		if err == nil {
//			w.Write(respData)
//		}
//		putAgent(pa)
//	}()
//	if handleFun, ok := agentRpcHandlers[arg.GetMsgID()]; ok {
//		im := GetMeta(arg.GetMsgID())
//		if im == nil {
//			log.Infof("Service NotFound MsgID=%d", arg.GetMsgID())
//			resp.Code = pb.ErrCode_ServiceNotFound
//			return
//		}
//		arg2, err := im.DecodeArg(arg.Payload)
//		if err != nil {
//			resp.Code = pb.ErrCode_InternalErr
//			return
//		}
//		reply, herr := handleFun(pa, arg2)
//		if herr != nil {
//			resp.Code = herr.ECode()
//			return
//		}
//		payload, err := im.EncodeReply(reply)
//		if err != nil {
//			resp.Code = pb.ErrCode_InternalErr
//		} else {
//			resp.Code = pb.ErrCode_Success
//			resp.Payload = payload
//		}
//	} else {
//		resp.Code = NoRpcHandleErr.ECode()
//	}
//}

// OnRouterRpcCall req: *pb.ReqArg  resp: *pb.RespMsg
func OnRouterRpcCall(arg *pb.ReqArg) (*pb.RespMsg, errcode2.IErr) {

	pa := getAgent(arg.SessionId)
	defer func() {
		putAgent(pa)
	}()

	var resp = new(pb.RespMsg)
	resp.MsgID = arg.MsgID

	pa.Uid = arg.Uid

	if handleFun, ok := agentRpcHandlers[arg.GetMsgID().ID()]; ok {
		im, err := meta.GetMeta(arg.GetMsgID().ID())
		if err != nil {
			log.Infof("Service NotFound MsgID=%d", arg.GetMsgID())
			return nil, pb.ErrCode_MetaCoderNotFound
		}

		arg2, err := im.DecodeArg(arg.Payload)
		if err != nil {
			return nil, errcode2.WithErrcode(-33, err)
		}

		reply, herr := handleFun(pa, arg2)
		if herr != nil {
			resp.Code = herr.ECode()
			return resp, nil
		}

		payload, err := im.EncodeReply(reply)
		if err != nil {
			return nil, errcode2.WithErrcode(-44, err)
		}

		resp.Code = pb.ErrCode_Success
		resp.Payload = payload
		return resp, nil
	}

	resp.Code = pb.ErrCode_ProtoMsgIdNoHandles
	return resp, nil
}

// CallFunc req: *pb.Request  resp: *pb.Response
func CallFunc(uid string, arg *pb.Request, handle AgentRpcHandler) *pb.Response {

	pa := getAgent(uid)
	defer func() {
		putAgent(pa)
	}()

	var resp = new(pb.Response)
	resp.MsgID = arg.MsgID

	pa.Uid = uid

	im, err := meta.GetMeta(arg.GetMsgID().ID())
	if err != nil {
		log.Infof("Service NotFound MsgID=%d", arg.GetMsgID())
		resp.Code = pb.ErrCode_MetaCoderNotFound
		return resp
	}

	arg2, err := im.DecodeArg(arg.Payload)
	if err != nil {
		log.Infof("MsgID=%d req arg decode err: %v", arg.GetMsgID(), err)
		resp.Code = pb.ErrCode_InternalErr
		return resp
	}

	reply, herr := handle(pa, arg2)
	if herr != nil {
		resp.Code = herr.ECode()
		return resp
	}

	payload, err := im.EncodeReply(reply)
	if err != nil {
		log.Infof("MsgID=%d resp arg encode err: %v", arg.GetMsgID(), err)
		resp.Code = pb.ErrCode_InternalErr
		return resp
	}

	resp.Code = pb.ErrCode_Success
	resp.Payload = payload
	return resp
}

func RespFun(uid string, arg *pb.Response, handle AgentRespHandler) errcode2.IErr {

	pa := getAgent(uid)
	defer func() {
		putAgent(pa)
	}()

	pa.Uid = uid

	im, err := meta.GetMeta(arg.GetMsgID().ID())
	if err != nil {
		log.Infof("Service NotFound MsgID=%d", arg.GetMsgID())
		return pb.ErrCode_MetaCoderNotFound
	}

	arg2, err := im.DecodeReply(arg.Payload)
	if err != nil {
		log.Infof("MsgID=%d req arg decode err: %v", arg.GetMsgID(), err)
		return pb.ErrCode_InternalErr
	}

	return handle(pa, arg2)
}
