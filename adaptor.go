package log

type Adaptor interface {
	Log(*Entry)
}
