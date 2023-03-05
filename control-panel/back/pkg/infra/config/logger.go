package config

import "go.uber.org/zap"

var Logger *zap.Logger

type LogField zap.Field

func init() {
	Logger, _ := zap.NewProduction()
	defer func() {
		_ = Logger.Sync()
	}()
}
