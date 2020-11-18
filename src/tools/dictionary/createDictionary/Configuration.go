package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
)

/*
	Configuration is a final internal state of the application after the
	command line arguments are parsed and processed.
*/

type Configuration struct {
	//
	// General Config.
	//
	Force bool
	Debug bool
	//
	// Log Configuration
	//
	Log logger.Configuration
	//
	// Reporting Configuration
	//
	InputFile  string // Text file source.
	OutputFile string // Dictionary file output.
}
