package common

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BuildLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	status := os.Getenv("STATUS")

	if status == "PRODUCTION" {
		config.OutputPaths = []string{
			"./log.txt",
		}

		config.ErrorOutputPaths = []string{
			"./log.txt",
		}
	} else {
		config.OutputPaths = []string{
			"./log_test.txt",
		}

		config.ErrorOutputPaths = []string{
			"./log_test.txt",
		}
	}

	logger, err := config.Build()

	if err != nil {
		log.Fatal("failed to create log file. err : " + err.Error())
	}

	return logger
}
