package core

import "go.uber.org/zap"

func CreateLogger(name string) *zap.SugaredLogger {
	return zap.S().Named(name)
}
