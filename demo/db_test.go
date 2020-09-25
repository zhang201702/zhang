package main

import (
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zlog"
	"testing"
)

func TestDB(t *testing.T) {

	db := z.GetDB("dcp")

	r, err := db.Table("v_user").FindAll()
	if err != nil {
		zlog.Error(err, "error")
	}
	list := r.ToList()
	zlog.Debug(list)
}
