package log

import (
	"context"
	"os"
	"runtime"
)

type Logger struct {
	level     Level
	appenders []Appender
	encoder   Encoder
}

func NewLogger(appenders ...Appender) *Logger {
	return &Logger{
		level:     LevelTrace,
		appenders: appenders,
		encoder:   &defaultEncoder{},
	}
}

func (logger *Logger) With(keyvals ...interface{}) *Logger {
	ctxLogger := *logger
	ctxLogger.encoder = newContextEncoder(logger.encoder, keyvals)
	return &ctxLogger
}

func (logger *Logger) Tag(tag string) *Logger {
	tagLogger := *logger
	tagLogger.encoder = newTagEncoder(logger.encoder, tag)
	return &tagLogger
}

func (logger *Logger) GetLevel() Level {
	return logger.level
}

func (logger *Logger) SetLevel(level Level) {
	logger.level = level
}

func (logger *Logger) Enabled(level Level) bool {
	return level >= logger.level
}

func (logger *Logger) Trace(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelTrace) {
		logger.output(ctx, LevelTrace, msg, keyvals)
	}
}

func (logger *Logger) Debug(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelDebug) {
		logger.output(ctx, LevelDebug, msg, keyvals)
	}
}

func (logger *Logger) Info(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelInfo) {
		logger.output(ctx, LevelInfo, msg, keyvals)
	}
}

func (logger *Logger) Warning(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelWarning) {
		logger.output(ctx, LevelWarning, msg, keyvals)
	}
}

func (logger *Logger) Error(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelError) {
		logger.output(ctx, LevelError, msg, keyvals)
	}
}

func (logger *Logger) Fatal(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelError) {
		logger.output(ctx, LevelFatal, msg, keyvals)
	}
	os.Exit(-1)
}

func (logger *Logger) output(ctx context.Context, level Level, msg string, keyvals []interface{}) {
	entry := logger.encoder.Encode(ctx, level, msg, keyvals)

	var ok bool
	_, entry.File, entry.Line, ok = runtime.Caller(2)
	if !ok {
		entry.File = "???"
		entry.Line = -1
	}

	for _, appender := range logger.appenders {
		appender.Append(entry)
	}
}
