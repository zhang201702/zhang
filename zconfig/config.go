package zconfig

import (
	"github.com/zhang201702/zhang/zfile"
	"github.com/zhang201702/zhang/zlog"
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
	err := zfile.OpenJson("./config.json", &Config)
	if err != nil {
		zlog.LogError(err, "zconfig.init", "读取config.json 异常", err)
	}
	Config = ConfigInfo{}
}

var defaultConfig *map[string]interface{}

func Default() *map[string]interface{} {
	if defaultConfig == nil {
		defaultConfig = new(map[string]interface{})
		//defaultConfig = make(map[string]interface{},0)Z
		if err := zfile.OpenJson("./config.json", defaultConfig); err != nil {
			zlog.LogError(err, "zconfig.Default", "读取config.json 异常", err)
		}
	}
	return defaultConfig
}
