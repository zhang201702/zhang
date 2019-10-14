package zlog

import "log"

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
