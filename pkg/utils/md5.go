package utils

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

// CheckMD5 用于判断一个字符串是否是合法的MD5字符串
func CheckMD5(str string) bool {
	// 正则表达式匹配32个十六进制字符
	re := regexp.MustCompile(`^[a-f0-9]{32}$`)
	return re.MatchString(str)
}
