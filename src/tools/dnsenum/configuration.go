package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
)

/*
	Configuration is a final internal state of the application after the
	command line arguments are parsed and processed.
*/

type Configuration struct {
	Force          bool
	Debug          bool
	Concurrency    int
	Delay          int
	Depth          int
	DnsServer      string
	DnsRecordTypes string
	Domain         string
	Log            logger.Configuration
	Source         source.Configuration
	Output         string
	Timeout        int
}
