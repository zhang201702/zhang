package main

import (
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zlog"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	redis := z.Redis("abc")
	err := redis.Set("aaaa", "aaaaaaaaaaaaaaaaaa", time.Second*10)
	if err != nil {
		zlog.LogError(err, "err")
		t.Fail()
	}
	data, err := redis.Get("aaaa")
	t.Log(data, err)
	if err != nil {
		t.Fail()
	}

}
