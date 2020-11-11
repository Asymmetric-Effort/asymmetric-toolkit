package main

import (
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
)

/*
	Configuration is a final internal state of the application after the
	command line arguments are parsed and processed.
*/

type Configuration struct {
	Concurrency    int
	delay          int
	depth          int
	dictionary     string
	dnsServer      string
	dnsRecordTypes string
	domain         string
	log struct {
		level logLevel.LogLevel
		facility LogFacility.Facility
		target string
	}
	maxWordCount   int
	minWordSize    int
	maxWordSize    int
	output         string
	pattern        string
	source         string
	timeout        int
}
