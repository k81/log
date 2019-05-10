package log

import (
	"bytes"
	"fmt"

	"github.com/modern-go/gls"
)

var (
	PipeKVFormatter = &pipeKvFormatter{}
)

type pipeKvFormatter struct{}

func (l *pipeKvFormatter) Format(entry *Entry) (*bytes.Buffer, error) {
	buf := getBuffer()
	tag := "_undef"

	if entry.Tag != "" {
		tag = entry.Tag
	}

	fmt.Fprintf(buf, "[%s][%s][%d][%s:%d]%s||",
		entry.Level,
		entry.Time.Format("2006-01-02T15:04:05.000Z07:00"),
		gls.GoID(),
		entry.File,
		entry.Line,
		tag,
	)

	if entry.Msg != "" {
		buf.WriteString("msg=")
		buf.WriteString(entry.Msg)
		buf.WriteString("||")
	}

	for i := 0; i < len(entry.KeyVals); i += 2 {
		key, val := entry.KeyVals[i], entry.KeyVals[i+1]
		buf.WriteString(toString(key))
		buf.WriteString("=")
		fmt.Fprint(buf, val)
		buf.WriteString("||")
	}

	if len(entry.KeyVals) > 0 {
		buf.Truncate(buf.Len() - 2)
	}

	buf.WriteString("\n")

	return buf, nil
}
