package znet

import (
	"crypto/tls"
	"errors"
	"gopkg.in/gomail.v2"

	"github.com/zhang201702/zhang/zconfig"
	"github.com/zhang201702/zhang/zlog"
	"sync"
)

type EmailConf struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	FromName string `json:"formName"`
}

var emailConf = (*EmailConf)(nil)
var once = sync.Once{}

func SendEmail(toAddress, toName, subject, body string) error {

	once.Do(func() {
		temp := EmailConf{}
		if err := zconfig.Conf.GetStruct("email", &temp); err != nil {
			zlog.Log(err, "email.init", "未找找配置")
			return
		}
		emailConf = &temp
	})
	if emailConf == nil {
		return errors.New("未找到email的配置")
	}

	m := gomail.NewMessage()
	m.SetAddressHeader("From", emailConf.UserName, emailConf.FromName)
	m.SetAddressHeader("To", toAddress, toName)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(emailConf.Host, emailConf.Port, emailConf.UserName, emailConf.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
