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
	"github.com/jageros/hawos/registry"
	"github.com/jageros/hawos/transport"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	*transport.BaseServer
	svr      *grpc.Server
	register registry.Registrar
}

func New(ctx context.Context, opfs ...transport.SvrOpFn) *Server {
	s := &Server{
		BaseServer: transport.NewBaseServer(ctx, opfs...),
	}
	s.Options.Protocol = transport.GRPC

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

func (s *Server) Register(registrar registry.Registrar) error {
	if registrar == nil {
		return nil
	}
	s.register = registrar
	return s.register.Register(s.Ctx, s.Options.BuildServiceInstance())
}

func (s *Server) Deregister() {
	if s.register != nil {
		s.register.Deregister(s.Ctx, s.Options.BuildServiceInstance())
	}
}
