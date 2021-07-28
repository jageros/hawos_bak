/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    client
 * @Date:    2021/6/9 4:13 下午
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
	"github.com/jageros/hawos/errcode"
	"github.com/jageros/hawos/internal/conf"
	"github.com/jageros/hawos/log"
	"github.com/jageros/hawos/registry"
	"github.com/jageros/hawos/transport"
	"github.com/jageros/hawos/transport/rpc/resolver/discovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/connectivity"
	"sync"
)

type RpcFn func(cc *grpc.ClientConn)

type OpFn func(opt *Option)

type Client interface {
	DialWithFn(serviceName string, rpcFn RpcFn) errcode.IErr
}

type client struct {
	ctx    context.Context
	cancel context.CancelFunc
	d      registry.Discovery
	conns  map[string]*grpc.ClientConn // map[serverName]conn
	option *Option

	NameChan chan string
	rw       *sync.RWMutex
}

type Option struct {
	Id          string
	Name        string
	Type        transport.ProtoTy
	ServerNames []string
}

func (c *client) AddServerName(names ...string) {
	go func() {
		for _, name := range names {
			select {
			case <-c.ctx.Done():
				return
			case c.NameChan <- name:
			}

		}
	}()

}

func ServerNames(name ...string) OpFn {
	return func(opt *Option) {
		opt.ServerNames = name
	}
}

func NewClient(ctx context.Context, d registry.Discovery, opFns ...OpFn) *client {
	ctx2, cancel := context.WithCancel(ctx)
	cli := &client{
		ctx:    ctx2,
		cancel: cancel,
		d:      d,
		conns:  map[string]*grpc.ClientConn{},
		option: &Option{
			Type: transport.RpcClient,
		},
		rw:       &sync.RWMutex{},
		NameChan: make(chan string, 64),
	}

	for _, opf := range opFns {
		opf(cli.option)
	}
	return cli
}

func (c *client) Init(id, name string) error {
	c.option.Id = id
	c.option.Name = name
	return nil
}

func (c *client) getConn(name string) (*grpc.ClientConn, error) {
	c.rw.RLock()
	conn, ok := c.conns[name]
	c.rw.RUnlock()
	if ok {
		return conn, nil
	}

	c.rw.Lock()
	defer c.rw.Unlock()

	// 为避免别的协程已经创建，进行检测后再创建
	conn, ok = c.conns[name]
	if ok {
		return conn, nil
	}

	target := fmt.Sprintf("%s:///%s", discovery.Name, name)
	builder := discovery.NewBuilder(c.ctx, c.d)

	ctx, cancel := context.WithTimeout(c.ctx, conf.RPC_CALL_TIMEOUT)
	defer cancel()

	cc, err := grpc.DialContext(
		ctx,
		target,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy":"%s"}`, roundrobin.Name)), // This sets the initial balancing policy.
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithResolvers(builder),
	)
	if err != nil {
		return nil, err
	}
	c.conns[name] = cc

	return cc, nil
}

func (c *client) Serve() error {
	go func() {
		for _, name := range c.option.ServerNames {
			select {
			case <-c.ctx.Done():
				return
			case c.NameChan <- name:
			}
		}
	}()

	builder := discovery.NewBuilder(c.ctx, c.d)
	for {
		select {
		case <-c.ctx.Done():
			return c.ctx.Err()
		case name := <-c.NameChan:
			c.rw.RLock()
			_, ok := c.conns[name]
			c.rw.RUnlock()
			if ok {
				continue
			}
			target := fmt.Sprintf("%s:///%s", discovery.Name, name)
			cc, err := grpc.DialContext(
				c.ctx,
				target,
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy":"%s"}`, roundrobin.Name)), // This sets the initial balancing policy.
				grpc.WithInsecure(),
				grpc.WithBlock(),
				grpc.WithResolvers(builder),
			)
			if err != nil {
				return err
			}
			c.rw.Lock()
			c.conns[name] = cc
			c.rw.Unlock()
		}
	}
}
func (c *client) Stop() {
	for _, cc := range c.conns {
		cc.Close()
	}
	c.cancel()
}

func (c *client) Info() string {
	return fmt.Sprintf("ServerName=%s Type=%s ID=%s", c.option.Name, c.option.Type, c.option.Id)
}

// ---

func (c *client) DialWithFn(serviceName string, rpcFn RpcFn) errcode.IErr {
	cc, err := c.getConn(serviceName)
	if err != nil {
		log.Errorf("DialWithFn getConn err: %v", err)
		return errcode.WithErrcode(-11, err)
	}

	state := cc.GetState()
	if state != connectivity.Ready {
		ctx, _ := context.WithTimeout(c.ctx, conf.RPC_CALL_TIMEOUT)
		cc.WaitForStateChange(ctx, state)
	}
	if cc.GetState() == connectivity.Ready {
		rpcFn(cc)
		return nil
	} else {
		errMsg := fmt.Sprintf("%s Service Conn NotReady!", serviceName)
		log.Errorf(errMsg)
		return errcode.New(-22, errMsg)
	}
}
