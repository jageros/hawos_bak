/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    app
 * @Date:    2021/5/28 12:18 下午
 * @package: app
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package app

import (
	"context"
	"fmt"
	"github.com/jageros/hawos/log"
	mode2 "github.com/jageros/hawos/mode"
	registry2 "github.com/jageros/hawos/registry"
	transport2 "github.com/jageros/hawos/transport"
	uuid2 "github.com/jageros/hawos/uuid"
	"github.com/oklog/oklog/pkg/group"
	"github.com/oklog/run"
	"syscall"
)

type Application struct {
	options *Options
	ctx     context.Context
}

type OpFn func(options *Options)

type Options struct {
	appID    string
	appName  string
	mode     mode2.MODE
	svrs     []transport2.IServer
	register registry2.Registrar
}

func ID(id interface{}) OpFn {
	idStr := fmt.Sprintf("%v", id)
	return func(options *Options) {
		options.appID = idStr
	}
}

func Name(name string) OpFn {
	return func(options *Options) {
		options.appName = name
	}
}

func Mode(mode mode2.MODE) OpFn {
	return func(options *Options) {
		options.mode = mode
	}
}

func Servers(svrs ...transport2.IServer) OpFn {
	return func(options *Options) {
		options.svrs = svrs
	}
}

func Registry(r registry2.Registrar) OpFn {
	return func(options *Options) {
		options.register = r
	}
}

// New 创建应用
func New(ctx context.Context, opfs ...OpFn) *Application {
	ap := &Application{
		options: &Options{},
		ctx:     ctx,
	}
	for _, opf := range opfs {
		opf(ap.options)
	}

	if ap.options.appID == "" {
		var err error
		ap.options.appID, err = uuid2.NewNumStr("X-App-Id")
		if err != nil {
			log.Errorf("NewAppId NewNumStr Err: %v", err)
			ap.options.appID = uuid2.New()
		}
	}

	return ap
}

// Run 运行应用
func (a *Application) Run() {
	ctx, cancel := context.WithCancel(a.ctx)
	defer cancel()

	// 新建goroutine初始化并起服
	g := group.Group{}

	for _, svr := range a.options.svrs {
		s := svr

		g.Add(func() error {
			defer func() {
				log.Infof("Server Over: [ %s ]", s.Info())
			}()
			err := s.Init(a.options.appID, a.options.appName)
			if err != nil {
				return err
			}
			if ss, ok := s.(transport2.IRegistry); ok {
				err = ss.Register(a.options.register)
				if err != nil {
					return err
				}
			}

			return s.Serve()
		}, func(err error) {
			if ss, ok := s.(transport2.IRegistry); ok {
				ss.Deregister()
			}

			s.Stop()
		})

	}

	// 监听系统信号
	exFn, errFn := run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM)
	g.Add(func() error {
		return exFn()
	}, func(err error) {
		errFn(err)
	})

	log.Errorf("========= ApplicationStop: %v =========", g.Run())
}
