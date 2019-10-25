package zconfig

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zfile"
	"github.com/zhang201702/zhang/zlog"
)

var Debug = false
var Conf *gjson.Json

func init() {
	if Conf == nil {
		filePath := getDefaultConfigPath()
		if filePath != "" {
			m := z.Map{}
			if err := zfile.OpenJson(filePath, &m); err != nil {
				zlog.LogError(err, "zconfig.Default", "读取config.json 异常", err)
			}
			Conf = gjson.New(m)
			Debug = Conf.GetBool("Debug")
		}
	}
}

func getDefaultConfigPath() string {
	if zfile.PathExists("./config.json") {
		return "./config.json"
	} else if zfile.PathExists("./config/config.json") {
		return "./config/config.json"
	}
	return ""
}

func Get(key string) interface{} {
	return Conf.Get(key)
}
