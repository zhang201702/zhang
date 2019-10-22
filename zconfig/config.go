package zconfig

import (
	"errors"
	"github.com/gogf/gf/util/gconv"
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
	Debug     bool
}

type RabbitMQInfo struct {
	Addr     string
	UserName string
	Password string
	Port     string
}

var Config ConfigInfo
var Debug = false
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
	Debug = Config.Debug
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

func GetConfig(key string) interface{} {
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

func GetString(key string) (string,error) {
	info := GetConfig(key)
	if info == nil {
		return gconv.String(info), nil
	}
	return "",errors.New("未找到配置")
}

func GetInt(key string) (int,error) {
	info := GetConfig(key)
	if info == nil {
		return  gconv.Int(info),nil
	}
	return 0,errors.New("未找到配置")
}

func GetFloat(key string) (float64,error) {
	info := GetConfig(key)
	if info == nil {
		return  gconv.Float64(info),nil
	}
	return 0,errors.New("未找到配置")
}

func GetMap(key string) (z.Map,error) {
	info := GetConfig(key)
	if info == nil {
		if r,ok := info.(map[string]interface{}); ok {
			return nil, errors.New("数据类型不匹配")
		}else{
			return r, nil
		}
	}
	return nil,errors.New("未找到配置")
}