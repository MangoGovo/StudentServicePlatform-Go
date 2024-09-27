package utils

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// IsValidEmail 使用正则表达式来判断邮箱地址是否合法
func IsValidEmail(email string) bool {
	// 定义一个常见的邮箱正则表达式
	// 注意：这个正则表达式并不是完美的，但它可以覆盖大多数常见的邮箱格式
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
func GenerateVerifyCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	// 定义验证码的字符集
	var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var sb strings.Builder
	for i := 0; i < length; i++ {
		// 随机选择一个字符
		index := rand.Intn(len(chars))
		sb.WriteRune(chars[index])
	}
	return sb.String()
}
