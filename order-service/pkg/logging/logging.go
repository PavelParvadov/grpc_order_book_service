package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var l *zap.Logger

func NewLogger() *zap.Logger {
	return l
}

func init() {

	logger := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "msg",
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,
		ConsoleSeparator: " | ",
	})

	core := zapcore.NewCore(logger, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	l = zap.New(core, zap.AddCaller())
}
