package log

import (
	"bytes"
	"sync"
)

const MaxCachedBufferSize = 1 << 16 // 64KB

var bufferPool *sync.Pool

func init() {
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
}

func getBuffer() *bytes.Buffer {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func putBuffer(buf *bytes.Buffer) {
	if buf == nil {
		return
	}

	if buf.Cap() > MaxCachedBufferSize {
		return
	}

	bufferPool.Put(buf)
}
