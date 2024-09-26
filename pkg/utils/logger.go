package utils

import (
	"io"
	"log"
	"os"
)

var Log *log.Logger

func InitLogger() {
	/*
		初始化日志，既输出到控制台，又输出到文件
	*/
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	//defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	Log = log.New(mw, "", log.LstdFlags)
	Log.Printf("%s", "123")
}
