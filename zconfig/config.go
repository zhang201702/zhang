package zconfig

import (
	"errors"
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"github.com/zhang201702/zhang/zcrypto"
	"github.com/zhang201702/zhang/zfile"
	"github.com/zhang201702/zhang/zlog"
	"log"
	"testing"
)

var IsInit = false
var Debug = true
var innerConfig map[string]interface{}
var Conf *gjson.Json
var CryptoKey = []byte("zhang67890123456")
var CryptoVi = []byte("1234567890123456")
var stopWatch = make(chan bool)
var configPath = "config/config.json"
var watching = false
var AppKey = "JCz2HK0r!tLmAAKcV3S^Rm#AMQB7iBBN"

func init() {
	testing.Init()
	var filePath string
	flag.StringVar(&filePath, "c", "", "配置文件路径: -c config.json")
	flag.Parse()

	//filePath, _ = gfile.Search(filePath)
	if filePath == "" {
		filePath = getDefaultConfigPath()
	}
	initDefault(filePath)

}

func initDefault(filePath string) {
	IsInit = true
	innerConfig = make(map[string]interface{})

	Conf = gjson.New(innerConfig)
	if filePath != "" {

		zlog.Log("配置信息", filePath)

		if err := zfile.OpenJson(filePath, &innerConfig); err != nil {
			zlog.LogError(err, "zconfig.Default", "读取config.json 异常", err)
		}
		Conf = gjson.New(innerConfig)
		Debug = Conf.GetBool("debug")
		zlog.IsDebug = Debug
		configPath = filePath
		//if watching {
		//	stopWatch <- true
		//}
		//go watchConfig()
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

func DecodeConfig(data string) string {
	return zcrypto.Decode(AppKey, "", data)
}
func watchConfig() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					SetDefaultPath(configPath)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Add(configPath)
	if err != nil {
		log.Fatal(err)
	}
	watching = true
	select {
	case <-stopWatch:
		return
	}
}
