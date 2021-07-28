/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    rpc
 * @Date:    2021/6/8 3:59 下午
 * @package: rpc
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package rpc

import (
	"context"
	"fmt"
	registry2 "github.com/jageros/hawos/registry"
	transport2 "github.com/jageros/hawos/transport"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	*transport2.BaseServer
	svr      *grpc.Server
	register registry2.Registrar
}

func New(ctx context.Context, opfs ...transport2.SvrOpFn) *Server {
	s := &Server{
		BaseServer: transport2.NewBaseServer(ctx, opfs...),
	}
	s.Options.Protocol = transport2.GRPC

	s.svr = grpc.NewServer()

	return s
}

func (s *Server) RegistryService(registryFunc func(svr *grpc.Server)) {
	registryFunc(s.svr)
}

func (s *Server) Serve() error {
	s.PrintInfo()
	addr := fmt.Sprintf("%s:%d", s.Options.Ip, s.Options.Port)
	li, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.svr.Serve(li)
}

func (s *Server) Stop() {
	s.svr.GracefulStop()
}

func (s *Server) Register(registrar registry2.Registrar) error {
	s.register = registrar
	return s.register.Register(s.Ctx, s.Options.BuildServiceInstance())
}

func (s *Server) Deregister() {
	s.register.Deregister(s.Ctx, s.Options.BuildServiceInstance())
}
