package zconfig

import (
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zfile"
	"github.com/zhang201702/zhang/zlog"
	"strings"
)

type RedisInfo struct {
	Addr     string
	Port     string
	Password string
	DB       int
}
type DBInfo struct {
	Addr     string
	DBName   string
	UserName string
	Password string
	Port     string
}
type WebSocketInfo struct {
	Url string
}
type ConfigInfo struct {
	DB        DBInfo
	Redis     RedisInfo
	RabbitMQ  RabbitMQInfo
	DataPath  string
	Websocket WebSocketInfo
	Port      int
}

type RabbitMQInfo struct {
	Addr     string
	UserName string
	Password string
	Port     string
}

var Config ConfigInfo

func init() {

	filePath := getDefaultConfigPath()
	if filePath != "" {
		err := zfile.OpenJson(filePath, &Config)
		if err != nil {
			zlog.LogError(err, "zconfig.init", "读取config.json 异常", err)
		}
	} else {
		Config = ConfigInfo{}
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

var defaultConfig z.Map

func Default() z.Map {
	if defaultConfig == nil {
		filePath := getDefaultConfigPath()
		if filePath != "" {
			defaultConfig = z.Map{}
			if err := zfile.OpenJson(filePath, &defaultConfig); err != nil {
				zlog.LogError(err, "zconfig.Default", "读取config.json 异常", err)
			}
		}
	}
	return defaultConfig
}

func GetCconfig(key string) interface{} {
	keys := strings.Split(key, ".")
	var c interface{} = Default()
	var result interface{} = nil
	for _, key := range keys {
		switch c.(type) {
		case map[string]interface{}:
			{
				temp := c.(map[string]interface{})
				if p, ok := temp[key]; ok {
					c = p
					result = c
				} else {
					return nil
				}
			}
		case z.Map:
			{
				temp := c.(z.Map)
				if p, ok := temp[key]; ok {
					c = p
					result = c
				} else {
					return nil
				}
			}
		default:
			return nil
		}
	}
	return result
}
