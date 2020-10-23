package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/LogDestination"
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
)

func (o *Logger) Setup(config *cli.CommandLine) {
	/*
		Initialize the log destination file handle (fp)
	*/
	o.level.Set(config.Get("LogLevel"))
	o.facility.Set(defaultLoggerFacility)

	switch config.Log.Target {
	case LogDestination.Stdout:
		o.writer = o.logWriterStdOut
	case LogDestination.File:
		panic("not implemented (logger file writer)")
		//o.writer =  o.logWriterFile
	case LogDestination.Syslog:
		panic("not implemented (logger syslog writer")
		//o.writer =  o.LogWriterSyslog
	default:
		panic("Invalid log destination.")
	}
}
