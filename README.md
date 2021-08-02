# Hawos
一个可供快速开发业务逻辑的脚手架


## 环境要求
+ Linux/Darwin
+ golang 1.16.5
+ Redis v6.2.2
+ Nsq v1.2.0 / Kafka v2.8.0
+ etcd v3.5.0


## Example
### http server
```go

package main

import (
	"context"
	"fmt"
	http2 "net/http"

	"imserver/consts"

	"github.com/gin-gonic/gin"
	"github.com/jageros/hawos/app"
	"github.com/jageros/hawos/flags"
	"github.com/jageros/hawos/log"
	"github.com/jageros/hawos/transport/http"
)

var appName = "http-server"

func init() {
	// 解析命令行参数
	flags.Parse(appName)
}

func apiHandle(engine *gin.Engine) {
	r := engine.Group("/api")
	r.GET("/info", func(c *gin.Context) {
		c.String(http2.StatusOK, "hello world")
	})
	/*
		... ...
	*/
}

func main() {
	// 创建父context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 初始化log
	source := fmt.Sprintf("%s%d", appName, flags.Options.ID)
	log.Init(flags.LogPath(), flags.LogLevel, true, log.SetCaller(), log.SetSource(source), log.SetStdout())
	defer log.Sync()

	// 创建http服务
	httpSvr := http.New(ctx, flags.HttpOption)
	httpSvr.RegistryHandlers(apiHandle) // 注册http路由

	// 创建应用，并run
	app.New(
		ctx,
		app.ID(flags.Options.ID),
		app.Name(appName),
		app.Mode(flags.Mode),
		app.Servers(httpSvr),
	).Run()
}
```

### grpc server
```go

package main

import (
	"context"
	"fmt"
	"github.com/jageros/hawos/errcode"
	"github.com/jageros/hawos/protoc"

	_ "imserver/protos/meta"

	"github.com/jageros/hawos/app"
	"github.com/jageros/hawos/etcd"
	"github.com/jageros/hawos/etcd/registry"
	"github.com/jageros/hawos/flags"
	"github.com/jageros/hawos/log"
	"github.com/jageros/hawos/selector"
	"github.com/jageros/hawos/transport/rpc"
)

var appName = "grpc-server"

func init() {
	// 解析命令行参数
	flags.Parse(appName)
}

func main() {
	// 创建父context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 初始化log
	source := fmt.Sprintf("%s%d", appName, flags.Options.ID)
	log.Init(flags.LogPath(), flags.LogLevel,
		true,
		log.SetCaller(),
		log.SetStdout(),
		log.SetSource(source),
	)
	defer log.Sync()

	// 创建rpc服务
	rpcSvr := rpc.New(ctx, flags.RpcOption)
	rpcSvr.RegistryService(rpc.RegistryRpcServer) // 注册rpc服务

	// 创建etcd客户端， 用于服务注册与发现
	cli, err := etcd.New(flags.EtcdOption)
	if err != nil {
		log.Panicf("Create etcd client err=%v", err)
	}

	// 创建服务注册
	r := registry.New(cli)
	
	// 注册服务
	protoc.RegisterAgentRpcHandler(pb.MsgID(1), func(agent *protoc.Agent, arg interface{}) (interface{}, errcode.IErr) {
		/*
			... ...
		*/
		return nil, nil
	})
	
	// 用于proto协议id注册进etcd
	stor := selector.New(ctx, cli)

	// 创建应用，并run
	app.New(
		ctx,
		app.ID(flags.Options.ID),
		app.Name(appName),
		app.Mode(flags.Mode),
		app.Servers(rpcSvr, stor),
		app.Registry(r),
	).Run()
}
```

### websocket
```go

package main

import (
	"context"
	"fmt"
	
	"github.com/jageros/hawos/app"
	"github.com/jageros/hawos/flags"
	"github.com/jageros/hawos/log"
	"github.com/jageros/hawos/transport/ws"
)

var appName = "websocket-server"

func init() {
	// 解析命令行参数
	flags.Parse(appName)
}

func main() {
	// 创建父context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 初始化log
	source := fmt.Sprintf("%s%d", appName, flags.Options.ID)
	//log.Init("./logs/frontend.log", flags.LogLevel, true, log.SetStdout(), log.SetCaller(), log.SetSource(source))
	log.Init(flags.LogPath(), flags.LogLevel, true, log.SetCaller(), log.SetStdout(), log.SetSource(source))
	defer log.Sync()

	// 创建websocket服务
	wsSvr := ws.New(ctx, flags.WsOption)
	wsSvr.RegistryReadFunc(func(uid string, rData []byte, writer ws.ISession) {
		/*
			... ...
		*/
	}) // 注册通用的websocket处理中间件
	wsSvr.RegisterDisconnectFunc(func(uid string) {
		/*
			... ...
		*/
	})

	// 创建应用，并run
	app.New(
		ctx,
		app.ID(flags.Options.ID),
		app.Name(appName),
		app.Mode(flags.Mode),
		app.Servers(wsSvr),
	).Run()
}
```