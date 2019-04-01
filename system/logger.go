/**
 * @author Jose Nidhin
 */
package system

import (
	"os"
	"strings"

	"github.com/josnidhin/go-rest-api-example/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(appConfig *config.Config) *zap.Logger {
	level := logLevel(appConfig)
	initFields := map[string]interface{}{
		"pid": os.Getpid(),
	}

	hostname, err := os.Hostname()
	if err != nil {
		initFields["hostname"] = hostname
	}

	logger, err := zap.Config{
		Encoding:      "json",
		Level:         zap.NewAtomicLevelAt(level),
		Development:   appConfig.Log.Development,
		InitialFields: initFields,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}.Build()

	if err != nil {
		panic(err)
	}

	logger.Debug("Logger initialised")

	return logger
}

func logLevel(appConfig *config.Config) zapcore.Level {
	switch strings.ToLower(appConfig.Log.Level) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
