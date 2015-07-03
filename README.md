Logga adds more convenience around Golang's log and syslog packages, primarily by adding two capabilities:

* Multiplexes logging into multiple outputs, like syslog, stderr or stdout
* Uses "log priorities" (error, warning, info, debug, etc) borrowed from syslog for everything

Here's the example of how to use it:
```Go
import (
	lg "github.com/kontsevoy/logga"
	"log"
	"log/syslog"
)

func main() {
	// configure two logging targets: syslog and stdandard out:
	sysl   := lg.AddTarget(lg.TargetSyslog, syslog.LOG_INFO)
	writer := lg.AddTarget(lg.TargetStdout, syslog.LOG_INFO)
	
	// set global flags:
	lg.SetFlags(log.LstdFlags)
	
	// syslog flags:
	sysl.SetFlags(0)

  // write into both logs (syslog and stdout) with 3 different priorities:
	lg.Info("This is info\n")
	lg.Warning("This is warning\n")
	lg.Error("This is error\n")

	// configure standard log to write via logga:
	log.SetOutput(writer)
	log.Println("logang log output")
}
```
