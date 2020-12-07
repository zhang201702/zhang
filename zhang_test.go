package zhang

import (
	"github.com/zhang201702/zhang/zencrypt"
	"github.com/zhang201702/zhang/zencrypt/jasypt"
	"github.com/zhang201702/zhang/zlog"
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
	//zencrypt.Scrypt.EnCode("aaa")
	//content,err := zencrypt.PBE.Encrypt("wallet_salt", "Dctp@123456")
	jm := jasypt.Decryptor{jasypt.AlgoPBEWithMD5AndDES, "wallet_salt"}
	jm.Decrypt()
	if err != nil {
		zlog.Error(err)
	} else {
		zlog.Log(content)
	}

}
