package log

import (
	"bytes"
	golog "log"
	"log/syslog"
)

type SyslogWriter struct {
	BaseWriter
	writer    *syslog.Writer
	formatter *golog.Logger
	buffer    *bytes.Buffer
	priority  syslog.Priority
}

func NewSyslogWriter(p syslog.Priority, tag string) LoggaWriter {
	sw := &SyslogWriter{priority: p}
	sw.writer, _ = syslog.New(p, tag)
	sw.buffer = new(bytes.Buffer)
	sw.formatter = golog.New(sw.buffer, "", 0)
	return LoggaWriter(sw)
}

// Write implements io.Writer interface
func (w *SyslogWriter) Write(b []byte) (int, error) {
	w.WriteP(syslog.LOG_NOTICE, string(b))
	return 0, nil
}

// WriteP implements logga.LoggaWriter interface
func (w *SyslogWriter) WriteP(p syslog.Priority, s string) {
	s = w.formatHeader(s)
	if p <= w.priority {
		w.writer.Write([]byte(s))
	}
}

func (w *SyslogWriter) SetFlags(f int) {
	w.formatter.SetFlags(f)
}

func (w *SyslogWriter) formatHeader(s string) string {
	w.formatter.Print(s)
	defer w.buffer.Reset()
	return w.buffer.String()
}

func (w *SyslogWriter) SetPrefix(s string) {
	w.formatter.SetPrefix(s)
}
