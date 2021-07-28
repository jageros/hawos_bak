/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    plmxs
 * @Date:    2021/7/23 2:33 下午
 * @package: plmxs
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package plmxs

import (
	"github.com/gin-gonic/gin"
	gp "github.com/mcuadros/go-gin-prometheus"
)

func RegistryPrometheus(engine *gin.Engine) {
	p := gp.NewPrometheus("gin")
	p.Use(engine)
}
