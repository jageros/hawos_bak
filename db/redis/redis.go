/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    redis
 * @Date:    2021/7/19 10:18 上午
 * @package: RDB
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	db2 "github.com/jageros/hawos/db"
	"time"
)

var RDB *Redis

type Redis struct {
	*db2.Database
	*redis.ClusterClient
}

func Initialize(ctx context.Context, opts ...db2.OpFn) {
	RDB = &Redis{
		Database: db2.NewDatabase(ctx, opts...),
	}

	rcc := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    RDB.Opt.Addrs,
		Username: RDB.Opt.Username,
		Password: RDB.Opt.Password,
	})

	RDB.ClusterClient = rcc

	go func() {
		tk := time.NewTicker(RDB.Opt.WaitTime)
		for {
			select {
			case <-RDB.Ctx.Done():
				RDB.Close()
				return
			case <-tk.C:
				RDB.Ping(RDB.Ctx)
			}
		}
	}()
}

// ======== string or int =========

func GetString(key string) string {
	return RDB.Get(RDB.Ctx, key).String()
}

func SetString(key, value string) error {
	return RDB.Set(RDB.Ctx, key, value, 0).Err()
}

func GetInt(key string) (int, error) {
	return RDB.Get(RDB.Ctx, key).Int()
}

func SetInt(key string, value int) error {
	return RDB.Set(RDB.Ctx, key, value, 0).Err()
}

func GetInt64(key string) (int64, error) {
	return RDB.Get(RDB.Ctx, key).Int64()
}

func GetUint64(key string) (uint64, error) {
	return RDB.Get(RDB.Ctx, key).Uint64()
}

// =========== set ===============

func AddMembersToSet(key string, values ...interface{}) (interface{}, error) {
	var cmds = []interface{}{"SADD", key}
	cmds = append(cmds, values...)
	return RDB.Do(RDB.Ctx, cmds...).Result()
}

func GetAllMembersFromSet(key string) (interface{}, error) {
	return RDB.Do(RDB.Ctx, "SMEMBERS", key).Result()
}

func DelMembersInSet(key string, values ...interface{}) error {
	var cmds = []interface{}{"SREM", key}
	cmds = append(cmds, values...)
	return RDB.Do(RDB.Ctx, cmds...).Err()
}

func MembersCountOfSet(key string) (int64, error) {
	return RDB.Do(RDB.Ctx, "SCARD", key).Int64()
}

// =========== Hash ===============

type Encoder interface {
	Marshal() (map[string]string, error)
	Unmarshal(value map[string]string) error
}

func map2fields(v map[string]string) []interface{} {
	var values []interface{}
	for key, value := range v {
		values = append(values, key, value)
	}
	return values
}

func SetCache(key string, v Encoder) error {
	m, err := v.Marshal()
	if err != nil {
		return err
	}
	values := map2fields(m)
	return RDB.HMSet(RDB.Ctx, key, values...).Err()
}

func GetCache(key string, v Encoder) error {
	result, err := RDB.HGetAll(RDB.Ctx, key).Result()
	if err != nil {
		return err
	}
	return v.Unmarshal(result)
}

func LockExec(key string, f func(key string)) error {
	ctx, cancel := context.WithTimeout(RDB.Ctx, RDB.Opt.WaitTime)
	defer cancel()
	lockKey := key + "-lock"
	var ok bool
	for !ok {
		select {
		case <-ctx.Done():
			errMsg := fmt.Sprintf("%s; key=%s has lock", ctx.Err().Error(), key)
			return errors.New(errMsg)
		default:
			ok = RDB.SetNX(RDB.Ctx, lockKey, 1, RDB.Opt.WaitTime).Val()
		}
	}

	f(key)

	err := RDB.Del(RDB.Ctx, lockKey).Err()
	for err != nil {
		select {
		case <-ctx.Done():
			errMsg := fmt.Sprintf("%s; key=%s Del err=%v", ctx.Err().Error(), key, err)
			return errors.New(errMsg)
		default:
			err = RDB.Del(RDB.Ctx, lockKey).Err()
		}
	}
	return err
}

func Do(cmds ...interface{}) (interface{}, error) {
	return RDB.Do(RDB.Ctx, cmds...).Result()
}

func Del(key string) error {
	return RDB.Del(RDB.Ctx, key).Err()
}

func Incr(key string) (int64, error) {
	return RDB.Incr(RDB.Ctx, key).Result()
}

func Context() context.Context {
	return RDB.Ctx
}
