package zlog

import (
	"log"
)

var IsDebug = true
var IsInfo = true

func Log(args ...interface{}) {
	logArgs := make([]interface{}, 0)
	for _, v := range args {
		logArgs = append(logArgs, v, "==>")
	}
	log.Println(logArgs)
}
func LogError(err error, args ...interface{}) {
	args = append(args, err)
	Log(args...)
}

func Debug(args ...interface{}) {
	if IsDebug {
		args = append([]interface{}{"debug"}, args...)
		Log(args...)
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
