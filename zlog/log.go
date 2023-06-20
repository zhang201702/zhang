package zlog

import "github.com/gogf/gf/frame/g"

var IsDebug = true
var IsInfo = true

func Log(args ...interface{}) {
	g.Log().Info(args...)
}
func LogError(err error, args ...interface{}) {
	args = append(args, err)
	g.Log().Error(args...)
}

func Debug(args ...interface{}) {
	if IsDebug {
		args = append([]interface{}{"debug"}, args...)
		g.Log().Debug(args...)
	}
}

func Info(args ...interface{}) {
	if IsInfo {
		args = append([]interface{}{"info"}, args...)
		Log(args...)
	}
}

func Error(err error, args ...interface{}) {
	args = append([]interface{}{"error", err}, args...)
	Log(args...)
}
