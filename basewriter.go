package log

import (
	"io"
	"log/syslog"
)

type LoggaWriter interface {
	io.Writer
	WriteP(syslog.Priority, string)
	GetPriority() syslog.Priority
	SetFlags(int)
	SetPrefix(string)
}

type BaseWriter struct {
	priority syslog.Priority
}

func (b *BaseWriter) GetPriority() syslog.Priority {
	return b.priority
}
