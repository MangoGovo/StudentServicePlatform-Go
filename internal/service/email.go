package service

import (
	"StuService-Go/internal/dao"
	"StuService-Go/internal/global"
	"StuService-Go/pkg/utils"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"time"
)

func SaveVerifyCode(receiver string, code string) error {
	verifyCodeExpireMinute := global.Config.GetInt("email.verifyCodeExpireMinute")
	err := dao.RedisSetKeyVal(ctx, receiver, code, time.Minute*time.Duration(verifyCodeExpireMinute))
	return err

}

func GetVerifyCode(email string) (string, error) {
	return dao.RedisGetKeyVal(ctx, email)
}

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
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return SaveVerifyCode(receiver, code)
}
