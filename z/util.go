package z

import (
	"github.com/gogf/gf/util/gconv"
	"math"
)

// 获取数值，dName小数位，向下取值
func Floor(data interface{}, dNum int) float64 {
	f := gconv.Float64(data)
	r := math.Pow10(dNum)
	return math.Floor(f*r) / r
}
