package main

import (
	"log"
	"log/syslog"

	l "github.com/kontsevoy/log"
)

func main() {
	// use our own logger natively:
	//slw := l.AddTarget(l.TargetSyslog, syslog.LOG_INFO)
	//slw.SetFlags(0)
	writer := l.AddTarget(l.TargetStdout, syslog.LOG_ERR)
	l.SetFlags(log.LstdFlags)

	l.Info("This is info\n")
	l.Warning("This is warning\n")
	l.Error("This is error\n")
	l.Printf("This is printf %v -> %v -> %v\n\n", 1, 2, 3)

	// plug our logger into standard Golang log:
	log.SetOutput(writer)
	log.Println("golang log.Println")
}
