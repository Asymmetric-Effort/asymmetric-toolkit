package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
)

/*
	Configuration is a final internal state of the application
	after the command line arguments are parsed and processed.
*/

type Configuration struct {
	// General Config.
	Force          bool
	Debug          bool
	// Log Configuration
	Log            logger.Configuration
	// Source Configuration
	Source         source.Configuration
	// Attack Configuration
	Concurrency    int
	Delay          int
	Depth          int
	DnsServer      string
	DnsRecordTypes string
	Domain         string
	Timeout        int
	// Reporting Configuration
	Output         string
}
