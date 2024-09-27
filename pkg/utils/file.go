package utils

import (
	uuid "github.com/satori/go.uuid"
	"path/filepath"
)

// GetFileType 用于获取文件拓展名
func GetFileType(filename string) string {
	return filepath.Ext(filename)
}

// GetUUID 用于生成UUID
func GetUUID() string {
	return uuid.NewV1().String()
}
