package zredis

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/util/gconv"
	"time"
)

type Redis struct {
	*gredis.Redis
}

func (redis *Redis) Set(key interface{}, value string, dur time.Duration) error {

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
	_, err := redis.Do("SET", args)
	return err
}

func (redis *Redis) Get(key interface{}) (string, error) {
	result, err := redis.Do("GET", key)
	return gconv.String(result), err
}

func (redis *Redis) Key(patten string) ([]interface{}, error) {
	result, err := redis.Do("KEYS", patten)
	if err != nil {
		panic(err)
	}
	return result.([]interface{}), err
}
