package zredis

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/g"
	"github.com/zhang201702/zhang/zconfig"
)

var isInitRedis = false

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
			gredis.SetConfigByStr(link.(string), name)
		}
	case string:
		gredis.SetConfigByStr(redisInfo.(string))
	}
	isInitRedis = true
}

func GetRedis(name ...string) *Redis {

	if !isInitRedis {
		initRedis()
	}
	return &Redis{
		Redis: g.Redis(name...),
	}
}
