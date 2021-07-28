/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    flag
 * @Date:    2021/6/18 4:36 下午
 * @package: flag
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package flags

import (
	"flag"
	"fmt"
	db2 "github.com/jageros/hawos/db"
	etcd2 "github.com/jageros/hawos/etcd"
	"github.com/jageros/hawos/log"
	mode2 "github.com/jageros/hawos/mode"
	transport2 "github.com/jageros/hawos/transport"
	yaml2 "github.com/jageros/hawos/yaml"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strconv"
	"strings"
	"time"
)

type optFunc func(opt *Option)

var (
	Mode        = mode2.DebugMode
	LogLevel    = log.DebugLevel
	HttpOption  transport2.SvrOpFn
	RpcOption   transport2.SvrOpFn
	WsOption    transport2.SvrOpFn
	NsqOption   transport2.SvrOpFn
	KafkaOption transport2.SvrOpFn
	EtcdOption  etcd2.OpFn
	RedisOption db2.OpFn
	Options     *Option
)

type Option struct {
	ID           int
	AppName      string
	ModeStr      string
	HttpIp       string
	HttpPort     int
	RpcIp        string
	RpcPort      int
	WsIp         string
	WsPort       int
	EtcdAddrs    string
	RedisAddrs   string
	NsqAddrs     string
	KafkaAddrs   string
	FrontendAddr string
	Config       string
	LogDir       string

	EtcdUser    string
	EtcdPasswd  string
	RedisUser   string
	RedisPasswd string
	RedisDBNo   int
}

func vInt(elem, defaultVal, val int) int {
	if val != defaultVal {
		return val
	}
	return elem
}

func vString(elem, defaultVal, val string) string {
	if val != defaultVal {
		return val
	}
	return elem
}

func (op *Option) parseFromYaml(cfg *yaml2.Config) {
	if cfg != nil {
		if cfg.AppID != 0 {
			op.ID = cfg.AppID
		}
		// === http ===
		if cfg.Listen.HttpIp != "" {
			op.HttpIp = cfg.Listen.HttpIp
		}
		if cfg.Listen.HttpPort != 0 {
			op.HttpPort = cfg.Listen.HttpPort
		}
		// === rpc ===
		if cfg.Listen.RpcIp != "" {
			op.RpcIp = cfg.Listen.RpcIp
		}
		if cfg.Listen.RpcPort != 0 {
			op.RpcPort = cfg.Listen.RpcPort
		}
		// === websocket ===
		if cfg.Listen.WsIp != "" {
			op.WsIp = cfg.Listen.WsIp
		}
		if cfg.Listen.WsPort != 0 {
			op.WsPort = cfg.Listen.WsPort
		}
		// === etcd ===
		if len(cfg.Etcd.Addrs) > 0 {
			op.EtcdAddrs = strings.Join(cfg.Etcd.Addrs, ";")
		}
		if cfg.Etcd.User != "" {
			op.EtcdUser = cfg.Etcd.User
		}
		if cfg.Etcd.Password != "" {
			op.EtcdPasswd = cfg.Etcd.Password
		}
		// === redis ===
		if len(cfg.Redis.Addrs) > 0 {
			op.RedisAddrs = strings.Join(cfg.Redis.Addrs, ";")
		}
		if cfg.Redis.DB != 0 {
			op.RedisDBNo = cfg.Redis.DB
		}
		if cfg.Redis.User != "" {
			op.RedisUser = cfg.Redis.User
		}
		if cfg.Redis.Password != "" {
			op.RedisPasswd = cfg.Redis.Password
		}
		// === nsq ===
		if len(cfg.Nsq.Addrs) > 0 {
			op.NsqAddrs = strings.Join(cfg.Nsq.Addrs, ";")
		}
		// === kafka ===
		if len(cfg.Kafka.Addrs) > 0 {
			op.KafkaAddrs = strings.Join(cfg.Kafka.Addrs, ";")
		}
		// === frontend addr ===
		if cfg.Listen.FrontendAddr != "" {
			op.FrontendAddr = cfg.Listen.FrontendAddr
		}

	}
}

func defaultOption() *Option {
	op := &Option{
		ID:         1,
		ModeStr:    "debug",
		HttpIp:     "127.0.0.1",
		HttpPort:   8010,
		RpcIp:      "127.0.0.1",
		RpcPort:    8030,
		WsIp:       "127.0.0.1",
		WsPort:     8050,
		EtcdAddrs:  "127.0.0.1:2379",
		RedisAddrs: "127.0.0.1:7001;127.0.0.1:7002;127.0.0.1:7003;127.0.0.1:7004;127.0.0.1:7005;127.0.0.1:7006",
		NsqAddrs:   "127.0.0.1:4161",
		KafkaAddrs: "127.0.0.1:9092",
		LogDir:     "logs",
	}
	return op
}

func Parse(name string, opts ...optFunc) {
	dp := defaultOption()

	for _, optf := range opts {
		optf(dp)
	}

	var (
		path           = flag.String("config", dp.Config, "Config file path")
		logPath        = flag.String("log-dir", dp.LogDir, "Log dir")
		id             = flag.Int("id", dp.ID, "Application Id")
		modeStr        = flag.String("mode", dp.ModeStr, "Server mode, default: debug, optional：debug/test/release")
		httpListenIp   = flag.String("http-ip", dp.HttpIp, "Http listen ip")
		httpListenPort = flag.Int("http-port", dp.HttpPort, "HTTP server port")
		rpcListenIp    = flag.String("rpc-ip", dp.RpcIp, "Rpc listen ip")
		rpcListenPort  = flag.Int("rpc-port", dp.RpcPort, "Rpc server port")
		wsListenIp     = flag.String("ws-ip", dp.WsIp, "Websocket listen ip")
		wsListenPort   = flag.Int("ws-port", dp.WsPort, "Websocket server port")
		etcdAddr       = flag.String("etcd-addrs", dp.EtcdAddrs, "Etcd addrs, use ; split")
		redisAddr      = flag.String("redis-addrs", dp.RedisAddrs, "Redis addrs, use ; split")
		nsqAddr        = flag.String("nsq-addrs", dp.NsqAddrs, "NSQ addrs, use ; split")
		kafkaAddr      = flag.String("kafka-addrs", dp.KafkaAddrs, "kafka addrs, use ; split")
		frontendAddr   = flag.String("frontend-addr", dp.FrontendAddr, "frontend addr")
	)
	flag.Parse()

	Options = new(Option)
	*Options = *dp
	if *path != "" {
		cfg := yaml2.Parse(*path)
		Options.parseFromYaml(cfg)
	}

	Options.AppName = name
	Options.ID = vInt(Options.ID, dp.ID, *id)
	Options.ModeStr = vString(Options.ModeStr, dp.ModeStr, *modeStr)
	Options.HttpIp = vString(Options.HttpIp, dp.HttpIp, *httpListenIp)
	Options.HttpPort = vInt(Options.HttpPort, dp.HttpPort, *httpListenPort)
	Options.RpcIp = vString(Options.RpcIp, dp.RpcIp, *rpcListenIp)
	Options.RpcPort = vInt(Options.RpcPort, dp.RpcPort, *rpcListenPort)
	Options.WsIp = vString(Options.WsIp, dp.WsIp, *wsListenIp)
	Options.WsPort = vInt(Options.WsPort, dp.WsPort, *wsListenPort)
	Options.EtcdAddrs = vString(Options.EtcdAddrs, dp.EtcdAddrs, *etcdAddr)
	Options.RedisAddrs = vString(Options.RedisAddrs, dp.RedisAddrs, *redisAddr)
	Options.NsqAddrs = vString(Options.NsqAddrs, dp.NsqAddrs, *nsqAddr)
	Options.KafkaAddrs = vString(Options.KafkaAddrs, dp.KafkaAddrs, *kafkaAddr)
	Options.FrontendAddr = vString(Options.FrontendAddr, dp.FrontendAddr, *frontendAddr)
	Options.LogDir = vString(Options.LogDir, dp.LogDir, *logPath)

	if Options.FrontendAddr == "" {
		Options.FrontendAddr = fmt.Sprintf("127.0.0.1:%d", Options.WsPort)
	}

	switch *modeStr {
	case mode2.TestMode.String():
		Mode = mode2.TestMode
	case mode2.ReleaseMode.String():
		Mode = mode2.ReleaseMode
		LogLevel = log.InfoLevel
	}

	HttpOption = func(opt *transport2.Option) {
		opt.Ip = Options.HttpIp
		opt.Port = uint16(Options.HttpPort)
		opt.Mode = Mode
	}

	RpcOption = func(opt *transport2.Option) {
		opt.Ip = Options.RpcIp
		opt.Port = uint16(Options.RpcPort)
		opt.Mode = Mode
	}

	WsOption = func(opt *transport2.Option) {
		opt.Ip = Options.WsIp
		opt.Port = uint16(Options.WsPort)
		opt.ReadTimeout = time.Second * 120
		opt.WriteTimeout = time.Second * 120
		opt.Mode = Mode
	}

	etcdAddrs := strings.Split(Options.EtcdAddrs, ";")
	EtcdOption = func(config *clientv3.Config) {
		config.Endpoints = etcdAddrs
		config.Username = Options.EtcdUser
		config.Password = Options.EtcdPasswd
	}

	redisAddrs := strings.Split(Options.RedisAddrs, ";")
	RedisOption = func(opt *db2.Option) {
		opt.Type = db2.Redis
		opt.DBName = strconv.Itoa(Options.RedisDBNo)
		opt.Addrs = redisAddrs
		opt.Username = Options.RedisUser
		opt.Password = Options.RedisUser
	}

	nsqAddrs := strings.Split(Options.NsqAddrs, ";")
	NsqOption = func(opt *transport2.Option) {
		opt.Protocol = transport2.Nsq
		opt.Endpoints = nsqAddrs
	}

	kafkaAddrs := strings.Split(Options.KafkaAddrs, ";")
	KafkaOption = func(opt *transport2.Option) {
		opt.Protocol = transport2.Kafka
		opt.Endpoints = kafkaAddrs
	}
}

func LogPath() string {
	if strings.HasSuffix(Options.LogDir, "/") {
		return fmt.Sprintf("%s%s.log", Options.LogDir, Options.AppName)
	}
	return fmt.Sprintf("%s/%s.log", Options.LogDir, Options.AppName)
}
