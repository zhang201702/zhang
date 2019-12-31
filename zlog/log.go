package zlog

import (
	"github.com/zhang201702/zhang/zconfig"
	"log"
)

func Log(args ...interface{}) {
	logArgs := make([]interface{}, 0)
	for _, v := range args {
		logArgs = append(logArgs, v, "==>")
	}
	log.Println(logArgs)
}
func Debug(args ...interface{}) {
	if zconfig.Debug {
		args = append(args, "Debug")
		Log(args...)
	}
}

func LogError(err error, args ...interface{}) {
	args = append(args, err)
	Log(args...)
}
