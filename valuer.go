package log

import "context"

// A Valuer generates a log value. When passed to Context.With in a value
// element (odd indexes), it represents a dynamic value which is re-evaluated
// with each log event.
type Valuer func(context.Context) interface{}

// bindValues replaces all value elements (odd indexes) containing a Valuer
// with their generated value.
func bindValues(ctx context.Context, keyvals []interface{}) {
	for i := 1; i < len(keyvals); i += 2 {
		if v, ok := keyvals[i].(Valuer); ok {
			keyvals[i] = v(ctx)
		}
	}
}

// containsValuer returns true if any of the value elements (odd indexes)
// contain a Valuer.
func containsValuer(keyvals []interface{}) bool {
	for i := 1; i < len(keyvals); i += 2 {
		if _, ok := keyvals[i].(Valuer); ok {
			return true
		}
	}
	return false
}
