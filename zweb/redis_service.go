package zweb

import (
	"encoding/json"
	"errors"
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zlog"
	"github.com/zhang201702/zhang/zredis"
	"reflect"
	"time"
)

type RedisService struct {
	Redis *zredis.Redis
}

var errorNoFound = errors.New("未存数据")
var NilFound = errors.New("defaultFun未返回数据")

func (rs *RedisService) GetRidesData(key string, v interface{}) error {
	mData, err := rs.Redis.Get(key)
	if err != nil {
		zlog.Error(err, "获取缓存异常!")
		return err
	}
	if mData == "" {
		return errorNoFound
	}
	err2 := json.Unmarshal([]byte(mData), v)
	if err2 != nil {
		zlog.Error(err2, "缓存数据解析异常!")
		return err2
	}
	return nil
}
func (rs *RedisService) SetRedisData(key string, v interface{}, duration time.Duration) {
	rs.Redis.Set(key, z.String(v), duration)
}

func (rs *RedisService) GetData(key string, result interface{}, duration time.Duration, defaultFun func() interface{}) (err error) {
	err = rs.GetRidesData(key, result)
	if err == nil {
		return
	}
	defer func() {
		if err1 := recover(); err1 != nil {
			zlog.Error(nil, err1)
			err = errors.New("redisService.GetData异常")
		}
	}()
	dValue := defaultFun()

	switch dValue.(type) {
	case error:
		return dValue.(error)
	}
	rv := reflect.ValueOf(result)
	value1 := reflect.ValueOf(dValue)
	v := rv.Elem()
	v.Set(value1)
	//value1.Set(rv)
	//result = defaultFun()
	rs.SetRedisData(key, result, duration)
	//if value1.IsNil() {
	//	return NilFound
	//}

	return nil
}

func (rs *RedisService) Publish(channel string, data interface{}) {
	v, err := rs.Redis.Do("PUBLISH", channel, data)
	if err != nil {
		zlog.Error(err, "rs.Redis.publish", channel, data)
	}
	zlog.Info("rs.Redis.publish", v, channel, data)
}

func (rs *RedisService) Subscribe(channel string, receiver func(channel string, msg interface{})) {
	conn := rs.Redis.GetConn()
	_, err := conn.Do("SUBSCRIBE", channel)
	if err != nil {
		zlog.Error(err, "rs.Redis.Subscribe", channel)
	}
	go func() {
		for {
			msg, err := conn.Receive()
			if err != nil {
				zlog.Error(err, "rs.Redis.Subscribe.Receive", channel)
				continue
			}
			go receiver(channel, msg)
		}
	}()
}
