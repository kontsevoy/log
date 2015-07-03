package logga

import (
	"io"
	"log"
	"log/syslog"
)

type ConsoleWriter struct {
	BaseWriter
	logger *log.Logger
}

func NewConsoleWriter(p syslog.Priority, w io.Writer) LoggaWriter {
	return LoggaWriter(&ConsoleWriter{
		logger:     log.New(w, "", 0),
		BaseWriter: BaseWriter{priority: p},
	})
}

// Write implements io.Writer interface
func (w *ConsoleWriter) Write(b []byte) (int, error) {
	w.logger.Print(string(b))
	return len(b), nil
}

func (w *ConsoleWriter) WriteP(p syslog.Priority, s string) {
	w.logger.Print(s)
}

func (w *ConsoleWriter) SetFlags(f int) {
	w.logger.SetFlags(f)
}

func (w *ConsoleWriter) SetPrefix(s string) {
	w.logger.SetPrefix(s)
}
