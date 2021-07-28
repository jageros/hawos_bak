/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    uuid
 * @Date:    2021/5/28 6:45 下午
 * @package: uuid
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package uuid

import (
	"fmt"
	"github.com/jageros/hawos/consts"
	redis2 "github.com/jageros/hawos/db/redis"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strconv"
)

func NewNumStr(key string) (string, error) {
	key = "uuid-" + key
	num, err := redis2.Incr(key)
	if err != nil {
		return "", err
	}
	result := strconv.FormatInt(consts.BaseUid+num, 10)
	var uidStr string
	for _, n := range result {
		rn := rand.Intn(10)
		uidStr += fmt.Sprintf("%c%d", n, rn)
	}
	return uidStr, nil

}

func New() string {
	return uuid.NewV5(uuid.NewV4(), uuid.NewV1().String()).String()
}

func NewByStr(str string) string {
	return uuid.NewV5(uuid.NewV4(), str).String()
}
