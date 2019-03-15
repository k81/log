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
	return newTagLogger(logger, tag)
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
	logger.log(ctx, LevelTrace, "", msg, keyvals)
}

func (logger *Logger) Debug(ctx context.Context, msg string, keyvals ...interface{}) {
	logger.log(ctx, LevelDebug, "", msg, keyvals)
}

func (logger *Logger) Info(ctx context.Context, msg string, keyvals ...interface{}) {
	logger.log(ctx, LevelInfo, "", msg, keyvals)
}

func (logger *Logger) Warning(ctx context.Context, msg string, keyvals ...interface{}) {
	logger.log(ctx, LevelWarning, "", msg, keyvals)
}

func (logger *Logger) Error(ctx context.Context, msg string, keyvals ...interface{}) {
	logger.log(ctx, LevelError, "", msg, keyvals)
}

func (logger *Logger) Fatal(ctx context.Context, msg string, keyvals ...interface{}) {
	logger.log(ctx, LevelFatal, "", msg, keyvals)
	os.Exit(-1)
}

func (logger *Logger) log(ctx context.Context, level Level, tag, msg string, keyvals []interface{}) {
	if level < logger.level {
		return
	}

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
		Tag:     tag,
		Msg:     msg,
		KeyVals: keyvals,
	}

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
