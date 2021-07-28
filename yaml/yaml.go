/**
 * @Author: jager
 * @File: main
 * @Version: v1.0.0
 * @Date: 2021/4/30 10:29 上午
 * @package: test
 * @Description:
 *
 */

package yaml

import (
	"flag"
	"github.com/jageros/hawos/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	AppID int `yaml:"app_id"`
	Listen struct {
		HttpIp       string `yaml:"http_ip"`
		HttpPort     int    `yaml:"http_port"`
		RpcIp        string `yaml:"rpc_ip"`
		RpcPort      int    `yaml:"rpc_port"`
		WsIp         string `yaml:"ws_ip"`
		WsPort       int    `yaml:"ws_port"`
		FrontendAddr string `yaml:"frontend_addr"`
	}
	Etcd struct {
		Addrs    []string `yaml:"addrs"`
		User     string   `yaml:"user"`
		Password string   `yaml:"password"`
	}
	Redis struct {
		Addrs    []string `yaml:"addrs"`
		DB       int      `yaml:"db"`
		User     string   `yaml:"user"`
		Password string   `yaml:"password"`
	}
	Nsq struct {
		Addrs    []string `yaml:"addrs"`
		User     string   `yaml:"user"`
		Password string   `yaml:"password"`
	}
	Kafka struct {
		Addrs    []string `yaml:"addrs"`
		User     string   `yaml:"user"`
		Password string   `yaml:"password"`
	}
}

func Parse(path string) *Config {

	flag.Parse()
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Read Config yaml file err: %v", err)
	}

	conf := new(Config)
	err = yaml.Unmarshal(yamlFile, &conf)

	if err != nil {
		log.Fatalf("Read Config yaml Unmarshal err: %v", err)
	}
	return conf
}
