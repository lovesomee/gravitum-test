package logger

import "go.uber.org/zap"

// New creates a production-ready zap logger.
func New() (*zap.Logger, error) {
	return zap.NewProduction()
}
