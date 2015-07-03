package logga

import (
	"bytes"
	"log"
	"log/syslog"
)

type SyslogWriter struct {
	BaseWriter
	writer    *syslog.Writer
	formatter *log.Logger
	buffer    *bytes.Buffer
}

func NewSyslogWriter(p syslog.Priority, tag string) LoggaWriter {
	sw := &SyslogWriter{}
	sw.writer, _ = syslog.New(p, tag)
	sw.buffer = new(bytes.Buffer)
	sw.formatter = log.New(sw.buffer, "", 0)
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
	switch p {
	case syslog.LOG_INFO:
		w.writer.Info(s)
	case syslog.LOG_DEBUG:
		w.writer.Debug(s)
	case syslog.LOG_NOTICE:
		w.writer.Notice(s)
	case syslog.LOG_WARNING:
		w.writer.Warning(s)
	case syslog.LOG_ERR:
		w.writer.Err(s)
	case syslog.LOG_CRIT:
		w.writer.Crit(s)
	case syslog.LOG_ALERT:
		w.writer.Alert(s)
	default:
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
