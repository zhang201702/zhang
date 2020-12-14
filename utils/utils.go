package utils

import "github.com/gogf/gf/util/gconv"

func String(params ...interface{}) string {
	r := ""
	for i := range params {
		r += gconv.String(params[i])
	}
	return r
}
