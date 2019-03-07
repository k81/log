package log

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type StdAppender struct {
	sync.Mutex
	formatter Formatter
}

func NewStdAppender(formatter Formatter) *StdAppender {
	return &StdAppender{
		formatter: formatter,
	}
}

func (appender *StdAppender) Append(entry *Entry) {
	buf, err := appender.formatter.Format(entry)
	defer putBuffer(buf)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to format entry: reason=[%v]\n", err)
		return
	}

	appender.Lock()
	defer appender.Unlock()

	if _, err = io.Copy(os.Stderr, buf); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to copy data to writer: reason=[%v]\n", err)
		return
	}
}
