package utils

import (
	uuid "github.com/satori/go.uuid"
	"path/filepath"
)

func GetFileType(filename string) string {
	return filepath.Ext(filename)
}
func GetUUID() string {
	return uuid.NewV1().String()
}
