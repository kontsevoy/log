package logga

import (
	"fmt"
	"log/syslog"
	"os"
	"sync"
)

type LogTarget int
type LogFlags int

// possible LogTargets
const (
	TargetStdout = iota
	TargetStderr
	TargetSyslog
)

// possible LogFlags, compatible with Golang's log.SetFlags()
const (
	// Bits or'ed together to control what's printed. There is no control over the
	// order they appear (the order listed here) or the format they present (as
	// described in the comments).  A colon appears after these items:
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // the date: 2009/01/23
	Ltime                         // the time: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
	Ldebug                        // source file + line number of where log message originated
)

type Logga struct {
	writers []LoggaWriter
	wlock   sync.Mutex
	tag     string
	flags   int
}

var Std = NewLogga()

func NewLogga() *Logga {
	return &Logga{
		writers: make([]LoggaWriter, 0),
		tag:     os.Args[0],
	}
}

func Debug(s string, args ...interface{}) {
	Std.writeWithPriority(syslog.LOG_DEBUG, s, args...)
}

func Info(s string, args ...interface{}) {
	Std.writeWithPriority(syslog.LOG_INFO, s, args...)
}

func Notice(s string, args ...interface{}) {
	Std.writeWithPriority(syslog.LOG_NOTICE, s, args...)
}

func Warning(s string, args ...interface{}) {
	Std.writeWithPriority(syslog.LOG_WARNING, s, args...)
}

func Error(s string, args ...interface{}) {
	Std.writeWithPriority(syslog.LOG_ERR, s, args...)
}

func AddTarget(target LogTarget, p syslog.Priority) LoggaWriter {
	return Std.AddTarget(target, p)
}

func SetFlags(f int) {
	Std.SetFlags(f)
}

func SetPrefix(prefix string) {
	Std.SetPrefix(prefix)
}

func Prefix() string {
	return Std.Prefix()
}

func Print(v ...interface{}) {
	Notice(fmt.Sprint(v...))
}

func Println(v ...interface{}) {
	v = append(v, "\n")
	Notice(fmt.Sprint(v...))
}

func Printf(s string, v ...interface{}) {
	Notice(fmt.Sprintf(s, v...))
}

func (l *Logga) SetFlags(f int) {
	l.flags = f
	for _, w := range l.writers {
		w.SetFlags(f)
	}
}

func (l *Logga) SetPrefix(prefix string) {
	l.tag = prefix
	for _, w := range l.writers {
		w.SetPrefix(prefix)
	}
}

func (l *Logga) Prefix() string {
	return l.tag
}

func (l *Logga) AddTarget(target LogTarget, p syslog.Priority) LoggaWriter {
	l.wlock.Lock()
	defer l.wlock.Unlock()

	var w LoggaWriter

	switch target {
	case TargetSyslog:
		w = NewSyslogWriter(p, l.tag)
	case TargetStdout:
		w = NewConsoleWriter(p, os.Stdout)
	case TargetStderr:
		w = NewConsoleWriter(p, os.Stderr)
	default:
		panic("Unknown log target")
	}
	l.writers = append(l.writers, w)
	return w
}

func (l *Logga) Debug(s string, args ...interface{}) {
	l.writeWithPriority(syslog.LOG_DEBUG, s, args...)
}

func (l *Logga) Info(s string, args ...interface{}) {
	l.writeWithPriority(syslog.LOG_INFO, s, args...)
}

func (l *Logga) Warning(s string, args ...interface{}) {
	l.writeWithPriority(syslog.LOG_WARNING, s, args...)
}

func (l *Logga) Error(s string, args ...interface{}) {
	l.writeWithPriority(syslog.LOG_ERR, s, args...)
}

func (l *Logga) writeWithPriority(p syslog.Priority, s string, args ...interface{}) {
	l.wlock.Lock()
	defer l.wlock.Unlock()

	s = fmt.Sprintf(s, args...)
	for _, w := range l.writers {
		w.WriteP(p, s)
	}
}
