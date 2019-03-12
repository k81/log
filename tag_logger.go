package log

import (
	"context"
	"os"
)

type TagLogger struct {
	*Logger
	Tag string
}

func (logger *TagLogger) Trace(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelTrace) {
		entry := logger.newEntry(ctx, LevelTrace, msg, keyvals)
		logger.log(entry)
	}
}

func (logger *TagLogger) Debug(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelDebug) {
		entry := logger.newEntry(ctx, LevelDebug, msg, keyvals)
		logger.log(entry)
	}
}

func (logger *TagLogger) Info(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelInfo) {
		entry := logger.newEntry(ctx, LevelInfo, msg, keyvals)
		logger.log(entry)
	}
}

func (logger *TagLogger) Warning(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelWarning) {
		entry := logger.newEntry(ctx, LevelWarning, msg, keyvals)
		logger.log(entry)
	}
}

func (logger *TagLogger) Error(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelError) {
		entry := logger.newEntry(ctx, LevelError, msg, keyvals)
		logger.log(entry)
	}
}

func (logger *TagLogger) Fatal(ctx context.Context, msg string, keyvals ...interface{}) {
	if logger.Enabled(LevelFatal) {
		entry := logger.newEntry(ctx, LevelFatal, msg, keyvals)
		logger.log(entry)
	}
	os.Exit(-1)
}

func (logger *TagLogger) newEntry(ctx context.Context, level Level, msg string, keyvals []interface{}) *Entry {
	entry := logger.Logger.newEntry(ctx, level, msg, keyvals)
	entry.Tag = logger.Tag
	return entry
}
