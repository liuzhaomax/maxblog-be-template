package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	logrus.SetLevel(logrus.InfoLevel) // Trace << Debug << Info << Warning << Error << Fatal << Panic
	InitializeLogging("golog.txt")
}

func InitializeLogging(logFile string) {
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开日志文件失败: " + err.Error())
		panic(err)
	}
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	logrus.SetFormatter(&logrus.TextFormatter{})
}

type GormLogger struct{}

func (*GormLogger) Print(v ...interface{}) {
	fileName := "golog.txt"
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("打开日志文件失败: " + err.Error())
		panic(err)
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetOutput(io.MultiWriter(src, os.Stdout))
	if v[0] == "sql" {
		logger.WithFields(logrus.Fields{
			"module":  "gorm",
			"type":    "sql",
			"rows":    v[5],
			"src_ref": v[1],
			"values":  v[4],
		}).Print(v[3])
	}
	if v[0] == "log" {
		logger.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
