package utils

import (
	"io"
	"log"
	"os"
)

var Log *log.Logger

// InitLogger 初始化日志,即输出到控制台,又输出到文件
func InitLogger() {
	/*

	 */
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	Log = log.New(mw, "", log.LstdFlags)
}
