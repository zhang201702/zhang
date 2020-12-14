package zredis

import (
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/util/gconv"
	"time"
)

type Redis struct {
	*gredis.Redis
	err error
}

func (redis *Redis) Set(key interface{}, value string, dur time.Duration) error {
	if redis.err != nil {
		return redis.err
	}
	args := make([]interface{}, 2, 4)
	args[0] = key
	args[1] = value

	if dur > 0 {
		if dur < time.Second || dur%time.Second != 0 {

			args = append(args, "px", int64(dur/time.Millisecond))
		} else {
			args = append(args, "ex", int64(dur/time.Second))
		}
	}
	_, err := redis.Do("SET", args...)
	return err
}

func (redis *Redis) Get(key interface{}) (string, error) {
	if redis.err != nil {
		return "", redis.err
	}
	result, err := redis.Do("GET", key)
	return gconv.String(result), err
}

func (redis *Redis) Key(patten string) ([]interface{}, error) {
	if redis.err != nil {
		return nil, redis.err
	}
	result, err := redis.Do("KEYS", patten)
	if err != nil {
		panic(err)
	}
	return result.([]interface{}), err
}

func (redis *Redis) Do(command string, args ...interface{}) (interface{}, error) {
	if redis.err != nil {
		return nil, redis.err
	}
	return redis.Redis.Do(command, args...)
}
func (redis *Redis) DoVar(command string, args ...interface{}) (*gvar.Var, error) {
	if redis.err != nil {
		return nil, redis.err
	}
	return redis.Redis.DoVar(command, args...)
}
