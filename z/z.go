package z

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
)

type Map map[string]interface{}

type Result struct {
	gjson.Json
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Code   string      `json:code`
}

func NewResult(result interface{}, err error) *Result {
	if err != nil {
		return &Result{
			Status: false,
			Msg:    err.Error(),
			Data:   result,
		}
	}
	var m map[string]interface{}
	switch result.(type) {
	case map[string]interface{}:
		m = result.(map[string]interface{})
	case Map:
		m = result.(Map)
	default:
		m = Map{
			"status": true,
			"data":   result,
		}
	}
	r := &Result{
		Status: gconv.Bool(m["result"]),
		Msg:    gconv.String(m["msg"]),
		Code:   gconv.String(m["code"]),
		Data:   m["data"],
	}
	r.Json = *gjson.New(r.Data)
	return r
}
