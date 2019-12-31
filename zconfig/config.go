package zconfig

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zfile"
	"github.com/zhang201702/zhang/zlog"
	"path/filepath"
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
	path, _ := filepath.Abs("config.json")
	if zfile.PathExists(path) {
		return path
	}
	path, _ = filepath.Abs("config/config.json")
	if zfile.PathExists(path) {
		return path
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
