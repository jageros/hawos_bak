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
	"github.com/jageros/hawos/protoc"
	"github.com/jageros/hawos/protos/pbf"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) ReqCall(ctx context.Context, arg *pbf.ReqArg) (*pbf.RespMsg, error) {
	resp, err := protoc.OnRouterRpcCall(arg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func RegistryRpcServer(s *grpc.Server) {
	pbf.RegisterRouterServer(s, &server{})
}
