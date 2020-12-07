package zconfig

import (
	"errors"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"github.com/zhang201702/zhang/zfile"
	"github.com/zhang201702/zhang/zlog"
)

var IsInit = false
var Debug = true
var innerConfig map[string]interface{}
var Conf *gjson.Json
var CryptoKey = []byte("zhang67890123456")
var CryptoVi = []byte("1234567890123456")

func init() {
	filePath := getDefaultConfigPath()
	initDefault(filePath)

}

func initDefault(filePath string) {
	if filePath != "" {
		innerConfig := new(map[string]interface{})
		zlog.Log("配置信息", filePath)
		if err := zfile.OpenJson(filePath, &innerConfig); err != nil {
			zlog.LogError(err, "zconfig.Default", "读取config.json 异常", err)
		}
		Conf = gjson.New(innerConfig)
		Debug = Conf.GetBool("Debug")
		zlog.IsDebug = Debug
		IsInit = true
	} else {
		zlog.LogError(errors.New("未找到配置信息"))
		Conf = gjson.New(new(map[string]interface{}))
	}
}

func getDefaultConfigPath() (path string) {

	path, _ = gfile.Search("config.json")
	zlog.Debug("path1", path)
	if path != "" {
		return path
	}
	path, _ = gfile.Search("config/config.json")
	zlog.Debug("path2", path)
	if path != "" {
		return path
	}
	return ""
}
func Get(key string, def ...interface{}) interface{} {
	return Conf.Get(key, def...)
}

func SetDefaultPath(path string) {
	if zfile.PathExists(path) {
		initDefault(path)
	}
}

func AddConfig(newConfig map[string]interface{}) {
	for k, v := range newConfig {
		innerConfig[k] = v
	}
}
