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

func SetLevelByName(levelName string) {
	defaultLogger.SetLevel(LevelName(levelName).ToLevel())
}

func GetLevel() Level {
	return defaultLogger.GetLevel()
}

func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
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
	defaultLogger.log(ctx, LevelTrace, "", msg, keyvals)
}

func Debug(ctx context.Context, msg string, keyvals ...interface{}) {
	defaultLogger.log(ctx, LevelDebug, "", msg, keyvals)
}

func Info(ctx context.Context, msg string, keyvals ...interface{}) {
	defaultLogger.log(ctx, LevelInfo, "", msg, keyvals)
}

func Warning(ctx context.Context, msg string, keyvals ...interface{}) {
	defaultLogger.log(ctx, LevelWarning, "", msg, keyvals)
}

func Error(ctx context.Context, msg string, keyvals ...interface{}) {
	defaultLogger.log(ctx, LevelError, "", msg, keyvals)
}

func Fatal(ctx context.Context, msg string, keyvals ...interface{}) {
	defaultLogger.log(ctx, LevelFatal, "", msg, keyvals)
	os.Exit(-1)
}
