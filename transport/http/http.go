/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    http
 * @Date:    2021/5/28 2:44 下午
 * @package: http
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package http

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/jageros/hawos/mode"
	"github.com/jageros/hawos/transport"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*transport.BaseServer
	svr *http.Server
}

func New(ctx context.Context, opfs ...transport.SvrOpFn) *Server {
	s := &Server{
		BaseServer: transport.NewBaseServer(ctx, opfs...),
	}

	s.Options.Protocol = transport.HTTP

	if s.Options.Mode == mode.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	addr := fmt.Sprintf("%s:%d", s.Options.Ip, s.Options.Port)
	s.svr = &http.Server{
		Addr: addr,
		BaseContext: func(listener net.Listener) context.Context {
			return s.Ctx
		},
	}
	return s
}

func (s *Server) RegistryHandlers(registryFun ...func(engine *gin.Engine)) {
	engine := gin.New()
	gin.ForceConsoleColor()
	engine.Use(logger(), gin.Recovery(), cors.Default())
	for _, f := range registryFun {
		f(engine)
	}
	s.svr.Handler = engine
}

func (s *Server) Serve() error {
	s.PrintInfo()
	return s.svr.ListenAndServe()
}
func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(s.Ctx, s.Options.CloseTimeout)
	defer cancel()
	s.svr.Shutdown(ctx)
}
