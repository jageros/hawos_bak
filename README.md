# Hawos demo
一个可供快速开发业务逻辑的脚手架

## 环境要求
+ Linux/Darwin
+ golang 1.16.5
+ Redis v6.2.2
+ Nsq v1.2.0 / Kafka v2.8.0
+ etcd v3.5.0

## 部署
+ 根据环境要求安装相关环境
+ 修改相关服务配置（配置文件：config/config.yaml和tools/start.sh脚本中的启动参数）
+ 编译：make plat=linux/darwin arch=amd64/arm/arm64
+ 启服：sh tools/start.sh
+ 停服：sh tools/stop.sh

## 服务简介
### frontend
``提供websocket服务，管理客户端长链``
+ [CPU占用分析](https://sp.hawtech.cn/hawpic/images/svg/cpu_frontend.svg)
+ [内存占用分析](https://sp.hawtech.cn/hawpic/images/svg/mem_frontend.svg)

### config
``提供http服务, 通过api获取配置信息等等``
+ [CPU占用分析](https://sp.hawtech.cn/hawpic/images/svg/cpu_config.svg)
+ [内存占用分析](https://sp.hawtech.cn/hawpic/images/svg/mem_config.svg)

### chat
``逻辑服（聊天室），处理聊天信息``
+ [CPU占用分析](https://sp.hawtech.cn/hawpic/images/svg/cpu_chat.svg)
+ [内存占用分析](https://sp.hawtech.cn/hawpic/images/svg/mem_chat.svg)

### test-client
``测试客户端``

``param: -id=[1,2,3,4...] -cnt=客户端数量``

## 架构
### 说明
+ API，表示通过http提供api服务，根据业务需求去实现相应功能
+ frontend，与client保持长连接，进来的请求通过协议ID选择逻辑层的相应服务进行转发，监听Nsq或者kafka的消息，然后推向给相应的client
+ chat（聊天服），逻辑层的一个代表，根据业务需求增加其他服务
+ 逻辑层的服务中通过协议id注册的服务都会把appName和协议id注册进etcd，frontend从etcd发现，根据协议请求的id获取相应的服务名称
+ 所有grpc请求使用的都是基于appName轮循的负载均衡
+ api服根据需要可以通过grpc请求frontend和逻辑层的服务
+ 服务注册与发现中分两种，其中一种是服务名称和地址（用于grpc负载均衡），另一种是服务名称和协议Id（用于frontend转发请求）

### 架构图
![框架图](doc/static/frame.png)

## 目录说明 ``（暂定）``
```
cmd： 业务代码

config: 配置文件目录

protos: 协议定义

tools: 工具脚本，包括协议编译，停起服脚本等

internal/cache ： 缓存实现

internal/pkg ： 通用代码实现
├── app    // app接口定义以及run实现
│   └── app.go  
├── db  // 数据库相关
│   ├── db.go
│   ├── mongo
│   │   └── mongo.go
│   ├── mysql
│   │   └── mysql.go
│   └── redis
│       └── redis.go
├── errcode  // 自定义带错误码的错误类实现
│   └── errcode.go
├── etcd  // etcd相关
│   ├── etcd.go
│   ├── registry  // 服务注册与发现
│   │   ├── registry.go
│   │   ├── registry_test.go
│   │   ├── service.go
│   │   └── watcher.go
│   └── tls_config.go
├── flags  // 命令行解析封装
│   └── flags.go
├── internal  // 全局不常变动的配置
│   └── conf
│       └── conf.go
├── jwt  // jwt 认证
│   └── jwt.go
├── log  // 自定义log实现
│   ├── adapter.go
│   └── log.go
├── mode // debug,release，test模式定义
│   └── mode.go
├── pprof // 性能分析相关
│   └── pprof.go
├── protoc // 协议编解码相关
│   ├── agent.go
│   └── rpc.go
├── recover   // 调用不可预知错误函数的recover捕获定义 
│   └── recover.go
├── registry  // 服务注册与发现接口定义
│   └── registry.go
├── selector  // 注册服务中注册的协议id进etcd，frontend服通过协议id转发请求到相应服
│   ├── code.go
│   ├── dscv.go
│   └── selector.go
├── sensitive  // 脏词过滤
│   ├── filter.go
│   └── trie.go
├── transport  // 传输层封装， 包括http，websocket， rpc， kafka， nsq
│   ├── http
│   │   ├── http.go
│   │   ├── httpc
│   │   │   ├── header.go
│   │   │   └── httpc.go
│   │   ├── logger.go
│   │   └── utils.go
│   ├── kafka
│   │   ├── pull.go
│   │   └── push.go
│   ├── nsq
│   │   ├── consumer.go
│   │   └── producer.go
│   ├── rpc
│   │   ├── client.go
│   │   ├── grpc.go
│   │   ├── resolver
│   │   │   ├── direct
│   │   │   │   ├── builder.go
│   │   │   │   └── resolver.go
│   │   │   └── discovery
│   │   │       ├── builder.go
│   │   │       ├── resolver.go
│   │   │       └── resolver_test.go
│   │   └── server.go
│   ├── transport.go  // 定于服务接口以及基类实现
│   └── ws      // websocket
│       ├── session.go
│       ├── ws.go
│       └── wsc
│           └── session.go
├─── uuid  // 分布式唯一ID生成， 通过Redis自增
│    └── uuid.go
└─── yaml  // yaml配置文件解析
     └── yaml.go
```