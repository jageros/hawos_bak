package recover

import (
	"fmt"
	errcode2 "github.com/jageros/hawos/errcode"
	"github.com/jageros/hawos/log"
	"reflect"
	"runtime"
)

func CatchPanic(f func() error) (err error) {
	defer func() {
		err1 := recover()
		if err1 != nil {
			fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
			log.Errorf("%s panic: %v", fn, err1)
			err = errcode2.New(1, fmt.Sprintf("%v", err1))
		}
	}()

	err = f()
	return
}
