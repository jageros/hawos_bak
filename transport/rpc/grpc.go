/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    grpc
 * @Date:    2021/7/9 6:15 下午
 * @package: rpc
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package rpc

import (
	"context"
	protoc2 "github.com/jageros/hawos/protoc"
	"github.com/jageros/hawos/protos/pb"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) ReqCall(ctx context.Context, arg *pb.ReqArg) (*pb.RespMsg, error) {
	resp, err := protoc2.OnRouterRpcCall(arg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func RegistryRpcServer(s *grpc.Server) {
	pb.RegisterRouterServer(s, &server{})
}
