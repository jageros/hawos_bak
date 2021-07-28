/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    mode
 * @Date:    2021/6/18 4:11 下午
 * @package: mode
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package mode

type MODE string

const (
	DebugMode   MODE = "debug"
	TestMode    MODE = "test"
	ReleaseMode MODE = "release"
)

func (m MODE) String() string {
	return string(m)
}
