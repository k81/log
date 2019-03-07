package log

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	OpenFlag = os.O_CREATE | os.O_APPEND | os.O_WRONLY | syscall.O_DSYNC
	OpenPerm = 0644
)

type FileAppender struct {
	sync.Mutex
	disableLock bool // enable log line bigger than 4k
	levelMask   Level
	fileName    string
	file        *os.File
	formatter   Formatter
}

func NewFileAppender(levelMask Level, fileName string, formatter Formatter) (appender *FileAppender, err error) {
	appender = &FileAppender{
		levelMask: levelMask,
		fileName:  fileName,
		formatter: formatter,
	}
	if appender.file, err = os.OpenFile(fileName, OpenFlag, OpenPerm); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to open log file: filename=%v, reason=%v\n", fileName, err)
		return
	}
	appender.start()
	return
}

func (appender *FileAppender) DisableLock() {
	appender.disableLock = true
}

func (appender *FileAppender) Append(entry *Entry) {
	buf, err := appender.formatter.Format(entry)
	defer putBuffer(buf)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to format entry: reason=[%v]\n", err)
		return
	}

	if !appender.disableLock {
		appender.Lock()
		defer appender.Unlock()
	}

	if entry.Level&appender.levelMask != 0 {
		if _, err = io.Copy(appender.file, buf); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: failed to copy data to err file writer: reason=[%v]\n", err)
		}
	}
}

func (appender *FileAppender) Rotate() {
	if !appender.disableLock {
		appender.Lock()
		defer appender.Unlock()
	}

	file, err := os.OpenFile(appender.fileName, OpenFlag, OpenPerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to rotate log file \"%s\", reopen, reason=%v", appender.fileName, err)
		return
	}

	if err = syscall.Dup2(int(file.Fd()), int(appender.file.Fd())); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to rotate log file \"%s\", dup2(), reason=%v", appender.fileName, err)
		return
	}

	if err = file.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to close log file \"%s\", close(), reason=%v", appender.fileName, err)
		return
	}
}

func (appender *FileAppender) start() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, RotateSignal)

	go func() {
		for {
			sig := <-ch

			switch {
			case sig == RotateSignal:
				appender.Rotate()
			}
		}
	}()
}
