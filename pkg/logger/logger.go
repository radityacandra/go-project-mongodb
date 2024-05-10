package logger

import "go.uber.org/zap"

func LoadLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return logger
}
