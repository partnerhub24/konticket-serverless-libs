package logger

import (
	"github.com/partnerhub24/konticket-serverless-pkg/environment"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	env := environment.GetEnv()

	var config zap.Config

	if env.App.AppEnv == "prod" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, field ...zap.Field) {
	log.Info(message, field...)
}

func Debug(message string, field ...zap.Field) {
	log.Debug(message, field...)
}

func Error(message any, field ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), field...)
	case string:
		log.Error(v, field...)
	}
}
