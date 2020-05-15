package zconfig

import (
	"errors"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zfile"
	"github.com/zhang201702/zhang/zlog"
)

var IsInit = false
var Debug = true
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
		zlog.Log("配置信息", filePath)
		if err := zfile.OpenJson(filePath, &m); err != nil {
			zlog.LogError(err, "zconfig.Default", "读取config.json 异常", err)
		}
		Conf = gjson.New(m)
		Debug = Conf.GetBool("Debug")
		zlog.IsDebug = Debug
		IsInit = true
	} else {
		zlog.LogError(errors.New("未找到配置信息"))
		Conf = gjson.New(z.Map{})
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
