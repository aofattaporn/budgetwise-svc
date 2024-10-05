package log

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/goproject/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Debug(msg string)
	Debugf(template string, args ...interface{})
	Info(msg string)
	Infof(template string, args ...interface{})
	Warn(msg string)
	Warnf(template string, args ...interface{})
	Error(msg string)
	Errorf(template string, args ...interface{})
	Fatal(msg string)
	Fatalf(template string, args ...interface{})
	WithCtx(ctx context.Context) ILogger
	WithField(field Field) *Logger
	ServiceInfo(msg string)
	ServiceInfof(template string, args ...interface{})
	RepoInfof(template string, args ...interface{})
}

type Logger struct {
	logger *zap.SugaredLogger
}

type CorrelationId struct{}
type Field map[string]interface{}

func InitZapLogger(cfg configs.ILogConfig) (ILogger, error) {

	cores := []zapcore.Core{}

	level := zap.NewAtomicLevel()
	writeStdout := zapcore.Lock(os.Stdout)
	if err := level.UnmarshalText([]byte(cfg.ConsoleLevel())); err != nil {
		return nil, fmt.Errorf("failed to unmarshal console level text for stdout: %v", err)
	}

	core := zapcore.NewCore(getLogEncoder(cfg.ConsoleIsJson(), cfg.ConsoleColor()), writeStdout, level)
	cores = append(cores, core)

	return &Logger{
		logger: zap.New(zapcore.NewTee(cores...)).Sugar(),
	}, nil
}

func getLogEncoder(isJson bool, color bool) zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	config.MessageKey = "message"
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	if isJson {
		return zapcore.NewJSONEncoder(config)
	}
	if color {
		config.EncodeTime = customTimeEncoder
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	return zapcore.NewConsoleEncoder(config)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// Append the colored timestamp to the encoder
	enc.AppendString(fmt.Sprintf("\x1b[96;1m%s\x1b[0m", t.Format("2006-01-02 15:04:05.000")))
}
