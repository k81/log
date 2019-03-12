package log

import "context"

type ctxKey int

var (
	keyContext           = ctxKey(0xf7f7f7f7)
	rootContext *Context = newContext()
)

func getContext(ctx context.Context) *Context {
	if logCtx, ok := ctx.Value(keyContext).(*Context); ok {
		return logCtx
	}
	return rootContext
}

type Context struct {
	keyvals   []interface{}
	hasValuer bool
}

func newContext() *Context {
	return &Context{}
}

func (c *Context) With(keyvals ...interface{}) *Context {
	if len(keyvals) == 0 {
		return c
	}
	kvs := append(c.keyvals, keyvals...)
	if len(kvs)%2 != 0 {
		kvs = append(kvs, ErrMissingValue)
	}
	return &Context{
		keyvals:   kvs[:len(kvs):len(kvs)],
		hasValuer: c.hasValuer || containsValuer(keyvals),
	}
}

func (c *Context) WithPrefix(keyvals ...interface{}) *Context {
	if len(keyvals) == 0 {
		return c
	}
	n := len(c.keyvals) + len(keyvals)
	if len(keyvals)%2 != 0 {
		n++
	}
	kvs := make([]interface{}, 0, n)
	kvs = append(kvs, keyvals...)
	if len(kvs)%2 != 0 {
		kvs = append(kvs, ErrMissingValue)
	}
	kvs = append(kvs, c.keyvals...)
	return &Context{
		keyvals:   kvs[:len(kvs):len(kvs)],
		hasValuer: c.hasValuer || containsValuer(keyvals),
	}
}
