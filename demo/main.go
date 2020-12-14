package main

import (
	"github.com/zhang201702/zhang"
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zlog"
)

func main() {
	s := zhang.Default()
	z.OpenBrowse(z.GetUrl())
	r := z.Redis("abc")
	a, err := r.Do("GET", "test")
	zlog.LogError(err, a)
	s.Run()
}
