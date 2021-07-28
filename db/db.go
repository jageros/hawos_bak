/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    db
 * @Date:    2021/7/19 10:39 上午
 * @package: db
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package db

import (
	"context"
	"errors"
	"strings"
	"time"
)

type IDB interface {
	Initialize() error
	Stop()
}

type DBTYPE uint8

const (
	Unknown DBTYPE = iota
	MongoDB
	Redis
	Mysql
)

type OpFn func(opt *Option)

type Option struct {
	Type     DBTYPE   // 数据库类型
	Addrs    []string // 连接地址
	Username string   // 用户名
	Password string   // 密码
	DBName   string   // 数据库名
	MaxConn  int64    // 最大连接数
	WaitTime time.Duration
}

func defaultOptions() *Option {
	return &Option{
		Type:     Redis,
		Addrs:    []string{"127.0.0.1:6379"},
		DBName:   "0",
		MaxConn:  64,
		WaitTime: time.Second * 30,
	}
}

type Database struct {
	Ctx    context.Context
	Cancel context.CancelFunc
	Opt    *Option
}

func NewDatabase(ctx context.Context, opts ...OpFn) *Database {
	ctx_, cancel := context.WithCancel(ctx)
	db := &Database{
		Ctx:    ctx_,
		Cancel: cancel,
		Opt:    defaultOptions(),
	}

	for _, opf := range opts {
		opf(db.Opt)
	}

	return db
}

func (d *Database) Addr() string {
	return strings.Join(d.Opt.Addrs, ",")
}

func (d *Database) Initialize() error {
	return errors.New("not implement db interface")
}

func (d *Database) Stop() {}
