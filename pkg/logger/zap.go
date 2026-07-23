package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// Init menginisialisasi global Zap logger.
func Init(level string, environment string) error {
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder
	if environment == "production" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		zapLevel,
	)

	log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// Replace global logger
	zap.ReplaceGlobals(log)

	return nil
}

// GetLogger mengembalikan instance logger global.
func GetLogger() *zap.Logger {
	if log == nil {
		// Fallback: init dengan default jika belum diinisialisasi
		if err := Init("info", "development"); err != nil {
			panic("failed to initialize logger: " + err.Error())
		}
	}
	return log
}

// Sync mem-flush buffer log sebelum shutdown.
func Sync() error {
	if log != nil {
		return log.Sync()
	}
	return nil
}

// SugaredLogger returns a sugared logger for convenience.
func SugaredLogger() *zap.SugaredLogger {
	return GetLogger().Sugar()
}

// RequestLogFields membuat field standard untuk request logging.
func RequestLogFields(method, path string, statusCode int, latency time.Duration, clientIP string) []zap.Field {
	return []zap.Field{
		zap.String("http_method", method),
		zap.String("path", path),
		zap.Int("status_code", statusCode),
		zap.Duration("latency", latency),
		zap.String("client_ip", clientIP),
	}
}
