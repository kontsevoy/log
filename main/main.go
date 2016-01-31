package main

import (
	"log"
	"log/syslog"

	l "github.com/kontsevoy/log"
)

func main() {
	// turn on syslog:
	slw := l.AddTarget(l.TargetSyslog, syslog.LOG_DEBUG)
	slw.SetFlags(0)

	// turn on stderr:
	writer := l.AddTarget(l.TargetStderr, syslog.LOG_WARNING)
	l.SetFlags(log.Llongfile)

	l.Info("This is info\n")
	l.Warning("This is warning\n")
	l.Error("This is error\n")
	l.Printf("This is printf %v -> %v -> %v\n\n", 1, 2, 3)

	// plug our logger into standard Golang log:
	log.SetFlags(0)
	log.SetOutput(writer)
	log.Println("golang log.Println")
}
