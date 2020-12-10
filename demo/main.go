package main

import (
	"github.com/zhang201702/zhang"
	"github.com/zhang201702/zhang/z"
)

func main() {
	s := zhang.Default()

	z.OpenBrowse(z.GetUrl())
	s.Run()
}
