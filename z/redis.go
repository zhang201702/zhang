package z

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/g"
	"github.com/zhang201702/zhang/zconfig"
	"github.com/zhang201702/zhang/zredis"
)

func init() {
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
}

func Redis(name ...string) *zredis.Redis {
	return &zredis.Redis{
		Redis: g.Redis(name...),
	}
}
