package logger

import (
	"github.com/sirupsen/logrus"
)

func LogInfo(args ...interface{}) {
	logrus.Info(args...)
}

func LogInfoF(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func LogError(args ...interface{}) {
	logrus.Error(args...)
}

func LogErrorF(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func LogPrint(args ...interface{}) {
	logrus.Print(args...)
}

func LogPrintF(format string, args ...interface{}) {
	logrus.Printf(format, args...)
}
