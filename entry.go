package log

import (
	"context"
	"errors"
	"time"
)

var ErrMissingValue = errors.New("(MISSING)")

type Entry struct {
	Tag     string
	Msg     string
	File    string
	Line    int
	Level   Level
	Time    time.Time
	KeyVals []interface{}
}

type Encoder interface {
	Encode(ctx context.Context, level Level, msg string, keyvals []interface{}) *Entry
}

type defaultEncoder struct{}

func (*defaultEncoder) Encode(ctx context.Context, level Level, msg string, keyvals []interface{}) *Entry {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, ErrMissingValue)
	}

	entry := &Entry{
		Time:    time.Now(),
		Level:   level,
		Msg:     msg,
		KeyVals: keyvals,
	}

	return entry
}

type contextEncoder struct {
	encoder   Encoder
	keyvals   []interface{}
	hasValuer bool
}

func newContextEncoder(encoder Encoder, keyvals []interface{}) Encoder {
	if len(keyvals) == 0 {
		return encoder
	}

	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, ErrMissingValue)
	}

	return &contextEncoder{
		encoder:   encoder,
		keyvals:   keyvals[:len(keyvals):len(keyvals)],
		hasValuer: containsValuer(keyvals),
	}
}

func (e *contextEncoder) Encode(ctx context.Context, level Level, msg string, keyvals []interface{}) *Entry {
	kvs := make([]interface{}, 0, len(e.keyvals)+len(keyvals)+2)
	kvs = append(kvs, e.keyvals...)
	kvs = append(kvs, keyvals...)

	if e.hasValuer {
		bindValues(ctx, kvs[:len(e.keyvals)])
	}

	return e.encoder.Encode(ctx, level, msg, kvs)
}

type tagEncoder struct {
	encoder Encoder
	tag     string
}

func newTagEncoder(encoder Encoder, tag string) *tagEncoder {
	return &tagEncoder{
		tag:     tag,
		encoder: encoder,
	}
}

func (e *tagEncoder) Encode(ctx context.Context, level Level, msg string, keyvals []interface{}) *Entry {
	entry := e.encoder.Encode(ctx, level, msg, keyvals)
	entry.Tag = e.tag
	return entry
}
