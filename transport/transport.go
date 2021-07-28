/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    server
 * @Date:    2021/6/8 4:31 下午
 * @package: transport
 * @Version: v1.0.0
 *
 * @Description: 定义服务参数接口
 *
 */

package transport

import (
	"context"
	"fmt"
	"github.com/jageros/hawos/internal/pkg/internal/conf"
	"github.com/jageros/hawos/internal/pkg/log"
	mode2 "github.com/jageros/hawos/mode"
	registry2 "github.com/jageros/hawos/registry"
	"strings"
	"time"
)

type ProtoTy string

func (p ProtoTy) String() string {
	return string(p)
}

type SvrOpFn func(opt *Option)

const (
	TCP       ProtoTy = "tcp"
	HTTP      ProtoTy = "http"
	GRPC      ProtoTy = "grpc"
	WS        ProtoTy = "websocket"
	RpcClient ProtoTy = "rpc-client"
	Kafka     ProtoTy = "kafka"
	Nsq       ProtoTy = "nsq"
	Selector  ProtoTy = "selector"
	Register  ProtoTy = "register"
	Discover  ProtoTy = "discover"
)

const (
	defaultIp           = "127.0.0.1"
	defaultPort         = uint16(8888)
	defaultReadTimeout  = conf.READ_TIMEOUT
	defaultWriteTimeout = conf.WRITE_TIMEOUT
	closeTimeout        = conf.CLOSE_TIMEOUT
)

type IServer interface {
	Init(id, name string) error
	Serve() error
	Stop()
	Info() string
}

type IRegistry interface {
	Register(registrar registry2.Registrar) error
	Deregister()
}

// ========================= BaseServer ==========================

type BaseServer struct {
	Ctx     context.Context
	Cancel  context.CancelFunc
	Options *Option
}

func NewBaseServer(ctx_ context.Context, opts ...SvrOpFn) *BaseServer {
	ctx, cancel := context.WithCancel(ctx_)
	bs := &BaseServer{
		Ctx:     ctx,
		Cancel:  cancel,
		Options: DefaultOptions(),
	}
	for _, opt := range opts {
		opt(bs.Options)
	}

	if len(bs.Options.Endpoints) <= 0 {
		bs.Options.Endpoints = []string{fmt.Sprintf("%s:%d", bs.Options.Ip, bs.Options.Port)}
	}

	return bs
}

func (bs *BaseServer) Init(id, name string) error {
	bs.Options.ID = id
	bs.Options.Name = name
	return nil
}

func (bs *BaseServer) Serve() error {
	<-bs.Ctx.Done()
	return bs.Ctx.Err()
}
func (bs *BaseServer) Stop() {
	bs.Cancel()
}
func (bs *BaseServer) Info() string {
	return fmt.Sprintf("Name=%s Type=%s Port=%d", bs.Options.Name, bs.Options.Protocol, bs.Options.Port)
}

func (bs *BaseServer) PrintInfo() {
	log.Infof("Listen IP=%s Port=%d Protocol=%s Mode=%s AppName=%s", bs.Options.Ip, bs.Options.Port, bs.Options.Protocol, bs.Options.Mode, bs.Options.Name)
}

func (bs *BaseServer) Addrs() string {
	return bs.Options.endpoints()
}

// ==================== End ========================

// =================== Option =====================

type Option struct {
	ID           string
	Name         string
	Ip           string
	Port         uint16 // 端口最大值：65535
	Protocol     ProtoTy
	Mode         mode2.MODE
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	CloseTimeout time.Duration
	Endpoints    []string
}

func (op *Option) endpoints() string {
	l := len(op.Endpoints)
	if l <= 0 {
		return fmt.Sprintf("%s:%d", op.Ip, op.Port)
	}

	return strings.Join(op.Endpoints, ";")
}

func (op *Option) BuildServiceInstance() *registry2.ServiceInstance {
	return &registry2.ServiceInstance{
		ID:       op.ID,
		Name:     op.Name,
		Type:     op.Protocol.String(),
		Version:  "v1",
		Endpoint: op.endpoints(),
	}
}

func DefaultOptions() *Option {
	return &Option{
		Ip:           defaultIp,
		Port:         defaultPort,
		Protocol:     TCP,
		Mode:         mode2.DebugMode,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		CloseTimeout: closeTimeout,
	}
}

func IP(ip string) SvrOpFn {
	return func(opt *Option) {
		opt.Ip = ip
	}
}

func Port(port uint16) SvrOpFn {
	return func(opt *Option) {
		opt.Port = port
	}
}

func Mode(mode_ mode2.MODE) SvrOpFn {
	return func(opt *Option) {
		opt.Mode = mode_
	}
}

func ReadTimeout(t time.Duration) SvrOpFn {
	return func(opt *Option) {
		opt.ReadTimeout = t
	}
}

func WriteTimeout(t time.Duration) SvrOpFn {
	return func(opt *Option) {
		opt.WriteTimeout = t
	}
}

func Endpoint(addr ...string) SvrOpFn {
	return func(opt *Option) {
		opt.Endpoints = addr
	}
}

// ====================== End ======================
