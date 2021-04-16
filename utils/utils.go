package utils

import (
	"github.com/gogf/gf/util/gconv"
	"strings"
)

func String(params ...interface{}) string {
	r := ""
	for i := range params {
		r += gconv.String(params[i])
	}
	return r
}

func Join(sep string, params ...interface{}) string {
	r := make([]string, 0)
	for i := range params {
		a := gconv.String(params[i])
		r = append(r, a)
	}
	str := strings.Join(r, sep)
	return str
}
