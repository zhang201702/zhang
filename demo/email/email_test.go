package email

import (
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/znet"
	"testing"
)

func TestEmail(t *testing.T) {
	r := z.NewResult(nil, nil)
	r.Set("a", "aaaaa")
	r.Set("b", "bb")
	znet.SendEmail("wat2288@163.com", "zhang", "test", "test")

}
