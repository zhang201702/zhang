package zlog

import (
	"errors"
	"testing"
)

func TestLog(t *testing.T) {
	Log("test", "OK")
	t.Fail()
}

func TestLogError(t *testing.T) {
	err := errors.New("tst error")
	LogError(err, "test", "error", "ok")
}
