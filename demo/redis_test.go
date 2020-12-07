package main

import (
	"github.com/zhang201702/zhang/z"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	redis := z.Redis()
	redis.Set("aaaa", "aaaaaaaaaaaaaaaaaa", time.Second*10)
	data, err := redis.Get("aaaa")
	t.Log(data, err)
	if err != nil {
		t.Fail()
	}

}
