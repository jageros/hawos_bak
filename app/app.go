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
	"github.com/jageros/hawos/mode"
	"github.com/jageros/hawos/registry"
	"github.com/jageros/hawos/transport"
	"github.com/jageros/hawos/uuid"
	"github.com/oklog/oklog/pkg/group"
)

type Application struct {
	options *Options
	ctx     context.Context
}

type OpFn func(options *Options)

type Options struct {
	appID    string
	appName  string
	mode     mode.MODE
	svrs     []transport.IServer
	register registry.Registrar
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

func Mode(mode mode.MODE) OpFn {
	return func(options *Options) {
		options.mode = mode
	}
}

func Servers(svrs ...transport.IServer) OpFn {
	return func(options *Options) {
		options.svrs = svrs
	}
}

func Registry(r registry.Registrar) OpFn {
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
		ap.options.appID, err = uuid.NewRandNumStr("X-App-Id")
		if err != nil {
			log.Errorf("NewAppId NewNumStr Err: %v", err)
			ap.options.appID = uuid.New()
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
			if ss, ok := s.(transport.IRegistry); ok {
				err = ss.Register(a.options.register)
				if err != nil {
					return err
				}
			}

			return s.Serve()
		}, func(err error) {
			if ss, ok := s.(transport.IRegistry); ok {
				ss.Deregister()
			}

			s.Stop()
		})

	}

	g.Add(func() error {
		<-ctx.Done()
		return ctx.Err()
	}, func(err error) {
		cancel()
	})

	log.Infof(">>>>>>>>> ApplicationStop: 【%v】", g.Run())
}
