package z

import (
	"github.com/zhang201702/zhang/zredis"
)

func Redis(name ...string) *zredis.Redis {
	return zredis.GetRedis(name...)
}
