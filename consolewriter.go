package log

import (
	"io"
	golog "log"
	"log/syslog"
)

type ConsoleWriter struct {
	BaseWriter
	logger *golog.Logger
}

func NewConsoleWriter(p syslog.Priority, w io.Writer) LoggaWriter {
	return LoggaWriter(&ConsoleWriter{
		logger:     golog.New(w, "", 0),
		BaseWriter: BaseWriter{priority: p},
	})
}

// Write implements io.Writer interface
func (w *ConsoleWriter) Write(b []byte) (int, error) {
	w.logger.Print(string(b))
	return len(b), nil
}

func (w *ConsoleWriter) WriteP(p syslog.Priority, s string) {
	if w.GetPriority() >= p {
		w.logger.Output(4, s)
	}
}

func (w *ConsoleWriter) SetFlags(f int) {
	w.logger.SetFlags(f)
}

func (w *ConsoleWriter) SetPrefix(s string) {
	w.logger.SetPrefix(s)
}
