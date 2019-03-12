package log

import (
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
