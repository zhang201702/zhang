package zredis

import (
	"errors"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/g"
	"github.com/zhang201702/zhang/utils"
	"github.com/zhang201702/zhang/zconfig"
	"github.com/zhang201702/zhang/zlog"
	"sync"
)

var once = sync.Once{}

func initRedis() {

	infoName := "redis"
	redisInfo := zconfig.Get(infoName)
	if redisInfo == nil {
		return
	}
	switch redisInfo.(type) {
	case map[string]interface{}:
		myMap := redisInfo.(map[string]interface{})
		for name, link := range myMap {
			err := gredis.SetConfigByStr(link.(string), name)
			zlog.Log("注册redis", name, err)
		}
	case string:
		err := gredis.SetConfigByStr(redisInfo.(string))
		zlog.Log("注册redis", redisInfo.(string), err)
	}
}

func GetRedis(name ...string) (result *Redis) {
	once.Do(func() {
		initRedis()
	})

	defer func() {

		if err := recover(); err != nil {
			err1 := errors.New(utils.String("创建redis异常,name:", name))
			switch err.(type) {
			case error:
				err1 = errors.New(utils.String(err1.Error(), ",info:", err.(error).Error()))
			case string:
				err1 = errors.New(utils.String(err1.Error(), ",info:", err.(string)))
			}
			zlog.Error(err1, "zredis.GetRedis 异常", name)
			result = &Redis{err: err1}
			return
		}
		zlog.Log("获取Redis 成功", name)
	}()
	result = &Redis{
		Redis: g.Redis(name...),
	}
	return result
}
