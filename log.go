package log

import (
	"context"
	"os"
)

var DefaultLogger *Logger = NewLogger(NewStdAppender(PipeKVFormatter))

func SetLevelByName(levelName LevelName) {
	DefaultLogger.SetLevel(levelName.ToLevel())
}

func GetLevel() Level {
	return DefaultLogger.GetLevel()
}

func Enabled(level Level) bool {
	return DefaultLogger.Enabled(level)
}

func With(keyvals ...interface{}) *Logger {
	if len(keyvals) == 0 {
		return DefaultLogger
	}

	return DefaultLogger.With(keyvals...)
}

func Tag(tag string) *Logger {
	return DefaultLogger.Tag(tag)
}

func Trace(ctx context.Context, msg string, keyvals ...interface{}) {
	DefaultLogger.Trace(ctx, msg, keyvals...)
}

func Debug(ctx context.Context, msg string, keyvals ...interface{}) {
	DefaultLogger.Debug(ctx, msg, keyvals...)
}

func Info(ctx context.Context, msg string, keyvals ...interface{}) {
	DefaultLogger.Info(ctx, msg, keyvals...)
}

func Warning(ctx context.Context, msg string, keyvals ...interface{}) {
	DefaultLogger.Warning(ctx, msg, keyvals...)
}

func Error(ctx context.Context, msg string, keyvals ...interface{}) {
	DefaultLogger.Error(ctx, msg, keyvals...)
}

func Fatal(ctx context.Context, msg string, keyvals ...interface{}) {
	DefaultLogger.Fatal(ctx, msg, keyvals...)
	os.Exit(-1)
}
