package logger

import (
	"testing"

	"github.com/razielsd/rzgrpcmock/server/internal/config"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateLogger Get initialized logger.
func CreateLogger(appCfg *config.Config) (*zap.Logger, error) {
	var logLevel zap.AtomicLevel
	err := logLevel.UnmarshalText([]byte(appCfg.LogLevel))
	if err != nil {
		return nil, err
	}
	cfg := zap.Config{
		Encoding:         "json",
		Level:            logLevel,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return log, nil
}

func TestLogger(t *testing.T) *zap.Logger {
	log, err := zap.NewDevelopment()
	require.NoError(t, err, "failed init test logger")
	return log
}
