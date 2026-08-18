package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(environment string) (*zap.Logger, error) {

	var cfg zap.Config

	if environment == "local" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	// Disable Debug and Info
	cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)

	cfg.Encoding = "console"

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return cfg.Build()
}
