package log

import (
	"context"
	"os"
)

type TagLogger struct {
	*Logger
	tag string
}

func newTagLogger(logger *Logger, tag string) *TagLogger {
	return &TagLogger{
		Logger: logger,
		tag:    tag,
	}
}

func (logger *TagLogger) Trace(ctx context.Context, msg string, keyvals ...interface{}) {
	logger.Logger.log(ctx, LevelTrace, logger.tag, msg, keyvals)
}

func (logger *TagLogger) Debug(ctx context.Context, msg string, keyvals ...interface{}) {
	logger.Logger.log(ctx, LevelDebug, logger.tag, msg, keyvals)
}

func (logger *TagLogger) Info(ctx context.Context, msg string, keyvals ...interface{}) {
	logger.Logger.log(ctx, LevelInfo, logger.tag, msg, keyvals)
}

func (logger *TagLogger) Warning(ctx context.Context, msg string, keyvals ...interface{}) {
	logger.Logger.log(ctx, LevelWarning, logger.tag, msg, keyvals)
}

func (logger *TagLogger) Error(ctx context.Context, msg string, keyvals ...interface{}) {
	logger.Logger.log(ctx, LevelError, logger.tag, msg, keyvals)
}

func (logger *TagLogger) Fatal(ctx context.Context, msg string, keyvals ...interface{}) {
	logger.Logger.log(ctx, LevelFatal, logger.tag, msg, keyvals)
	os.Exit(-1)
}
