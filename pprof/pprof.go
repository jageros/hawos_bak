/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    pprof
 * @Date:    2021/7/20 10:55 上午
 * @package: pprof
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package pprof

import (
	"fmt"
	"github.com/jageros/hawos/log"
	"os"
	"runtime/pprof"
)

var id_ int
var name_ string
var cf *os.File

func Start(id int, appName string) {
	id_ = id
	name_ = appName
	cpuProfile := fmt.Sprintf("pprof/cpu_%s%d.pprof", name_, id_)
	var err error
	cf, err = os.Create(cpuProfile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cf)
}

func Stop() {
	pprof.StopCPUProfile()
	cf.Close()

	memProfile := fmt.Sprintf("pprof/mem_%s%d.pprof", name_, id_)
	f, err := os.Create(memProfile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(f)
	f.Close()
}
