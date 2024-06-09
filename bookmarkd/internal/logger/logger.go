package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

var Logger *zap.Logger

var timeFormat = time.RFC3339

const service = "bookmarkd"

func InitializeLogger() {
	Logger, _ = zap.NewProduction()
}

func Info(m string) {
	t := time.Now().Format(timeFormat)
	Logger.Info(m, zap.String("service", service), zap.String("time", t))
}

func Infof(format string, params ...interface{}) {
	t := time.Now().Format(timeFormat)
	m := fmt.Sprintf(format, params...)
	Logger.Info(m, zap.String("service", service), zap.String("time", t))
}

func Error(m string) {
	t := time.Now().Format(timeFormat)
	Logger.Error(m, zap.String("service", service), zap.String("time", t))
}

func Errorf(format string, params ...interface{}) {
	t := time.Now().Format(timeFormat)
	m := fmt.Sprintf(format, params...)
	Logger.Error(m, zap.String("service", service), zap.String("time", t))
}

func Fatal(m string) {
	t := time.Now().Format(timeFormat)
	Logger.Fatal(m, zap.String("service", service), zap.String("time", t))
}

func Fatalf(format string, params ...interface{}) {
	t := time.Now().Format(timeFormat)
	m := fmt.Sprintf(format, params...)
	Logger.Fatal(m, zap.String("service", service), zap.String("time", t))
}
