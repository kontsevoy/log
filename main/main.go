package main

import (
	l "github.com/kontsevoy/log"
	"log"
	"log/syslog"
)

func main() {
	// use our own logger natively:
	slw := l.AddTarget(l.TargetSyslog, syslog.LOG_INFO)
	writer := l.AddTarget(l.TargetStdout, syslog.LOG_INFO)
	l.SetFlags(log.LstdFlags)
	slw.SetFlags(0)

	l.Info("This is info\n")
	l.Warning("This is warning\n")
	l.Error("This is error\n")

	l.Print(1, 2, 3)
	l.Printf("%v -> %v -> %v\n\n", 1, 2, 3)

	// plug our logger into standard Golang log:
	log.SetOutput(writer)
	//log.Println("--(golang log) ---> woo!")
}
