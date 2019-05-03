/**
 * @author Jose Nidhin
 */
package system

import (
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var loggerOnce sync.Once

func LoggerInstance(appConfig *Config, appVersion string) *zap.Logger {
	loggerOnce.Do(func() {
		level := logLevel(appConfig)
		initFields := map[string]interface{}{
			"appVersion": appVersion,
			"goVersion": runtime.Version(),
			"pid": os.Getpid(),
		}

		hostname, err := os.Hostname()
		if err != nil {
			initFields["hostname"] = hostname
		}

		logger, err = zap.Config{
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
				EncodeTime:     zapcore.ISO8601TimeEncoder,
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
	})

	return logger
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		logger.Info(
			"Request log",
			zap.String("method", r.Method),
			zap.String("url", r.RequestURI),
			zap.String("userAgent", r.Header.Get("User-Agent")),
			zap.String("httpVersion", r.Proto),
			zap.Duration("requestDuration", time.Since(start)),
		)
	})
}

func logLevel(appConfig *Config) zapcore.Level {
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
