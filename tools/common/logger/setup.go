package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/destination"
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
)

func (o *Logger) Setup(config *cli.Configuration) {
	/*
		Initialize the log destination file handle (fp)
	*/
	o.level.Set(config.Log.Level)
	o.facility.Set(defaultLoggerFacility)

	switch config.Log.Destination {
	case destination.Stdout:
		o.writer = o.logWriterStdOut
	case destination.File:
		panic("not implemented (logger file writer)")
		//o.writer =  o.logWriterFile
	case destination.Syslog:
		panic("not implemented (logger syslog writer")
		//o.writer =  o.LogWriterSyslog
	default:
		panic("Invalid log destination.")
	}
	o.Info("Logger::Setup() complete.")
}
