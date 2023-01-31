package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(provideLogger),
)

func provideLogger() (*zap.Logger, *zap.SugaredLogger) {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()

	return logger, slogger
}
