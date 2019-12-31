package zhang

import (
	"github.com/zhang201702/zhang/zencrypt"
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	s := Default()
	timer := time.NewTimer(time.Second * 5)
	go func() {
		for range timer.C {
			s.Shutdown()
		}
	}()
	s.Run()
}

func TestEncrypt(t *testing.T) {
	zencrypt.Scrypt.EnCode("aaa")
}
