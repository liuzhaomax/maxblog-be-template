package conf

import (
	"github.com/sirupsen/logrus"
	"io"
	"maxblog-be-template/internal/core"
	"os"
)

func init() {
	logrus.SetLevel(logrus.InfoLevel) // Trace << Debug << Info << Warning << Error << Fatal << Panic
	InitializeLogging("golog.txt")    // TODO 根据时间创建不同的日志文件，减小IO开支
}

func InitializeLogging(logFile string) {
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"失败方法": core.GetFuncName(),
		}).Panic(core.FormatError(902, err).Error())
	}
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	logrus.SetFormatter(&logrus.TextFormatter{})
}

type GormLogger struct{}

func (*GormLogger) Print(v ...interface{}) {
	fileName := "golog.txt"
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"失败方法": core.GetFuncName(),
		}).Panic(core.FormatError(902, err).Error())
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
