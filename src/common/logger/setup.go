package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/destination"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
)

func (o *Logger) Setup(config *cli.Configuration) {
	/*
		Initialize the log destination file handle (fp)
	*/
	o.Level.Set(config.Log.Level)
	o.Facility.Set(DefaultLoggerFacility)

	switch config.Log.Destination {
	case destination.Stdout:
		o.Writer = o.logWriterStdOut
	case destination.File:
		panic("not implemented (logger file writer)")
		//o.writer =  o.logWriterFile
	case destination.Syslog:
		panic("not implemented (logger syslog writer")
		//o.writer =  o.LogWriterSyslog
	default:
		panic("Invalid log destination.")
	}
}
