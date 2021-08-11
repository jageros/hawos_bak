/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    randslice
 * @Date:    2021/8/11 1:35 下午
 * @package: utils
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package utils

import "math/rand"

type element interface {
	Len() int
	Swap(i, j int)
}

func RandElements(its element) {
	l := its.Len()
	for i := 0; i < l; i++ {
		j := rand.Intn(l)
		its.Swap(i, j)
	}
}
