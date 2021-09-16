package logger

import (
	"go.uber.org/zap"
	"log"
)

var Logger *zap.Logger

func InitLogger(logLevel string) *zap.Logger {
	if Logger != nil {
		return Logger
	}
	var err error

	logConf := zap.NewProductionConfig()

	switch logLevel {
	case "debug":
		logConf.Level.SetLevel(zap.DebugLevel)
	case "info":
		logConf.Level.SetLevel(zap.InfoLevel)
	case "warn":
		logConf.Level.SetLevel(zap.WarnLevel)
	case "error":
		logConf.Level.SetLevel(zap.ErrorLevel)
	case "fatal":
		logConf.Level.SetLevel(zap.FatalLevel)
	default:
		logConf.Level.SetLevel(zap.InfoLevel)
	}

	Logger, err = logConf.Build()
	if err != nil {
		log.Fatalf("error initializing logger: %s", err)
	}
	return Logger

	return Logger
}

func GetLogger() *zap.Logger {
	return Logger
}
