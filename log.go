package log

import (
	"context"
	"os"
	"sync"
)

var (
	defaultLogger *Logger = NewLogger(NewStdAppender(PipeKVFormatter))
	tagLoggerPool         = &sync.Map{}
)

func SetLogger(logger *Logger) {
	defaultLogger = logger
}

func GetLogger() *Logger {
	return defaultLogger
}

func SetLevelByName(levelName LevelName) {
	defaultLogger.SetLevel(levelName.ToLevel())
}

func GetLevel() Level {
	return defaultLogger.GetLevel()
}

func Enabled(level Level) bool {
	return defaultLogger.Enabled(level)
}

func WithContext(ctx context.Context, keyvals ...interface{}) context.Context {
	if len(keyvals) == 0 {
		return ctx
	}

	logCtx := getContext(ctx).With(keyvals...)
	return context.WithValue(ctx, keyContext, logCtx)
}

func Tag(tag string) *TagLogger {
	logger, ok := tagLoggerPool.Load(tag)
	if ok {
		return logger.(*TagLogger)
	}

	tagLogger := defaultLogger.Tag(tag)
	tagLoggerPool.Store(tag, tagLogger)
	return tagLogger
}

func Trace(ctx context.Context, msg string, keyvals ...interface{}) {
	defaultLogger.Trace(ctx, msg, keyvals...)
}

func Debug(ctx context.Context, msg string, keyvals ...interface{}) {
	defaultLogger.Debug(ctx, msg, keyvals...)
}

func Info(ctx context.Context, msg string, keyvals ...interface{}) {
	defaultLogger.Info(ctx, msg, keyvals...)
}

func Warning(ctx context.Context, msg string, keyvals ...interface{}) {
	defaultLogger.Warning(ctx, msg, keyvals...)
}

func Error(ctx context.Context, msg string, keyvals ...interface{}) {
	defaultLogger.Error(ctx, msg, keyvals...)
}

func Fatal(ctx context.Context, msg string, keyvals ...interface{}) {
	defaultLogger.Fatal(ctx, msg, keyvals...)
	os.Exit(-1)
}
