package zlog

import (
	"log"
)

var IsDebug = true

func Log(args ...interface{}) {
	logArgs := make([]interface{}, 0)
	for _, v := range args {
		logArgs = append(logArgs, v, "==>")
	}
	log.Println(logArgs)
}
func Debug(args ...interface{}) {
	if IsDebug {
		args = append(args, "Debug")
		Log(args...)
	}
}

func LogError(err error, args ...interface{}) {
	args = append(args, err)
	Log(args...)
}
