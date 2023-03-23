/**
 * @author Jose Nidhin
 */
package logger

import (
	"os"
	"runtime"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	AppName, AppVersion string

	Level       string `json:"level"`
	Development bool   `json:"development"`
}

func New(logConfig *LogConfig) *zap.Logger {
	level := logLevel(logConfig.Level)

	initFields := map[string]interface{}{
		"appName":    logConfig.AppName,
		"appVersion": logConfig.AppVersion,
		"goVersion":  runtime.Version(),
		"pid":        os.Getpid(),
	}

	hostname, err := os.Hostname()
	if err != nil {
		initFields["hostname"] = hostname
	}

	logger, err := zap.Config{
		Encoding:      "json",
		Level:         zap.NewAtomicLevelAt(level),
		Development:   logConfig.Development,
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

	return logger
}

//
//func RequestLogger(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
//		start := time.Now()
//		next.ServeHTTP(res, req)
//
//		logger.Info(
//			"Request log",
//			zap.String("method", req.Method),
//			zap.String("url", req.RequestURI),
//			zap.String("userAgent", req.Header.Get("User-Agent")),
//			zap.String("httpVersion", req.Proto),
//			zap.Duration("requestDuration", time.Since(start)),
//		)
//	})
//}

func logLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
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
