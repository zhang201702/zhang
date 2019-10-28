package z

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
	"log"
	"reflect"
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

func NewResult(result interface{}, err error) *Result {
	if err != nil {
		return &Result{
			Status: false,
			Msg:    err.Error(),
			Data:   result,
			Err:    err,
		}
	}
	if result == nil {
		return &Result{
			Status: true,
			Msg:    "",
		}
	}
	var m map[string]interface{}
	tt := reflect.TypeOf(result)
	log.Println(tt.String())
	switch result.(type) {
	case map[string]interface{}:
		m = result.(map[string]interface{})
	case Map:
		m = result.(Map)
	case string:
		m = Map{
			"result": true,
			"msg":    result.(string),
		}
	default:
		m = Map{
			"result": true,
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
