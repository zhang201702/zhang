package z

import (
	"github.com/gogf/gf/encoding/gjson"
)

type Map map[string]interface{}

type Result struct {
	gjson.Json
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Code   string      `json:"code""`
	Err    error       `json:"_"`
}
