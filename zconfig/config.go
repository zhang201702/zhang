package zconfig

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zfile"
	"github.com/zhang201702/zhang/zlog"
)

var Debug = false
var Conf *gjson.Json
var CryptoKey = []byte("zhang67890123456")
var CryptoVi = []byte("1234567890123456")

func init() {
	filePath := getDefaultConfigPath()
	initDefault(filePath)
}

func initDefault(filePath string) {
	if filePath != "" {
		m := z.Map{}
		if err := zfile.OpenJson(filePath, &m); err != nil {
			zlog.LogError(err, "zconfig.Default", "读取config.json 异常", err)
		}
		Conf = gjson.New(m)
		Debug = Conf.GetBool("Debug")
	} else {
		Conf = gjson.New(z.Map{})
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
func SetDefaultPath(path string) {
	if zfile.PathExists(path) {
		initDefault(path)
	}
}
