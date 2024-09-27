package service

import (
	"StuService-Go/internal/global"
	"StuService-Go/pkg/utils"
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func SendVerifyCode(receiver string, code string) error {
	message := utils.GetEmailTemplate(code)

	host := global.Config.GetString("email.host")
	port := global.Config.GetInt("email.port")
	userName := global.Config.GetString("email.username")
	password := global.Config.GetString("email.password")

	m := gomail.NewMessage()
	m.SetHeader("From", userName)

	m.SetHeader("To", receiver)
	m.SetHeader("Subject", "学生服务平台注册验证码")

	m.SetBody("text/html", message)

	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d.DialAndSend(m)
}
