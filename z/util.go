package z

import (
  "github.com/gogf/gf/encoding/gjson"
  "github.com/gogf/gf/util/gconv"
  "github.com/gogf/gf/util/guid"
  "github.com/google/uuid"
  "github.com/zhang201702/zhang/utils"
  "math"
)

// 获取数值，dName小数位，向下取值
func Floor(data interface{}, dNum int) float64 {
  f := gconv.Float64(data)
  r := math.Pow10(dNum)
  return math.Floor(f*r) / r
}

func NewResult(result interface{}, err error) *Result {
  if err != nil {

    r := &Result{
      Status: false,
      Msg:    err.Error(),
      Data:   result,
      Err:    err,
    }
    r.Json = *gjson.New(r.Data)
    return r
  }

  if result == nil {
    r := &Result{
      Status: true,
      Msg:    "",
      Data:   make(map[string]interface{}),
    }
    r.Json = *gjson.New(r.Data)
    return r
  }
  var m map[string]interface{}
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

func (result *Result) GetFloor64(pattern string, dNum int) float64 {
  f := result.GetFloat64(pattern)
  return Floor(f, dNum)
}

func String(params ...interface{}) string {
  return utils.String(params...)
}

func NewMap(data interface{}) (result Map) {
  gj := gjson.New(data)
  result = gj.Map()
  return result
}

func UUID() string {
  guid.S()
  return uuid.New().String()
}
