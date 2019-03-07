package log

import "syscall"

type Appender interface {
	Append(*Entry)
}

type RotatableAppender interface {
	Appender
	Rotate()
}

var RotateSignal = syscall.SIGUSR1
