package log

import (
	"context"
	"os"
	"runtime"
	"time"
)

type Logger struct {
	level     Level
	appenders []Appender
}

func NewLogger(appenders ...Appender) *Logger {
	return &Logger{
		level:     LevelTrace,
		appenders: appenders,
	}
}

func (logger *Logger) Tag(tag string) *TagLogger {
	return &TagLogger{
		Logger: logger,
		Tag:    tag,
	}
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
		entry := logger.newEntry(ctx, LevelTrace, msg, keyvals)
		logger.log(entry)
	}
}

func (logger *Logger) Debug(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelDebug) {
		entry := logger.newEntry(ctx, LevelDebug, msg, keyvals)
		logger.log(entry)
	}
}

func (logger *Logger) Info(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelInfo) {
		entry := logger.newEntry(ctx, LevelInfo, msg, keyvals)
		logger.log(entry)
	}
}

func (logger *Logger) Warning(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelWarning) {
		entry := logger.newEntry(ctx, LevelWarning, msg, keyvals)
		logger.log(entry)
	}
}

func (logger *Logger) Error(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelError) {
		entry := logger.newEntry(ctx, LevelError, msg, keyvals)
		logger.log(entry)
	}
}

func (logger *Logger) Fatal(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelError) {
		entry := logger.newEntry(ctx, LevelFatal, msg, keyvals)
		logger.log(entry)
	}
	os.Exit(-1)
}

func (logger *Logger) newEntry(ctx context.Context, level Level, msg string, keyvals []interface{}) *Entry {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, ErrMissingValue)
	}

	logCtx := getContext(ctx)
	if len(logCtx.keyvals) > 0 {
		keyvals = append(logCtx.keyvals, keyvals...)
	}

	if logCtx.hasValuer {
		bindValues(ctx, keyvals[:len(logCtx.keyvals)])
	}

	entry := &Entry{
		Time:    time.Now(),
		Level:   level,
		Msg:     msg,
		KeyVals: keyvals,
	}

	return entry
}

func (logger *Logger) log(entry *Entry) {
	var ok bool
	_, entry.File, entry.Line, ok = runtime.Caller(3)
	if !ok {
		entry.File = "???"
		entry.Line = -1
	}

	for _, appender := range logger.appenders {
		appender.Append(entry)
	}
}
