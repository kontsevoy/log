package main

import (
	"github.com/kontsevoy/logga"
	"log"
	"log/syslog"
)

func main() {
	// use our own logger natively:
	slw := logga.AddTarget(logga.TargetSyslog, syslog.LOG_INFO)
	writer := logga.AddTarget(logga.TargetStdout, syslog.LOG_INFO)
	logga.SetFlags(log.LstdFlags)
	slw.SetFlags(0)

	logga.Info("This is info\n")
	logga.Warning("This is warning\n")
	logga.Error("This is error\n")

	logga.Print(1, 2, 3)
	logga.Printf("%v -> %v -> %v\n\n", 1, 2, 3)

	// plug our logger into standard Golang log:
	log.SetOutput(writer)
	//log.Println("--(golang log) ---> woo!")
}
