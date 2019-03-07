package log

import (
	"context"
	"os"
)

var defaultLogger *Logger = NewLogger(NewStdAppender(PipeKVFormatter))

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

func With(keyvals ...interface{}) *Logger {
	if len(keyvals) == 0 {
		return defaultLogger
	}

	return defaultLogger.With(keyvals...)
}

func Tag(tag string) *Logger {
	return defaultLogger.Tag(tag)
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
