package main

import "asymmetric-effort/asymmetric-toolkit/src/common/logger"

/*
	Configuration is a final internal state of the application after the
	command line arguments are parsed and processed.
*/

type Configuration struct {
	Force          bool
	Debug          bool
	Concurrency    int
	delay          int
	depth          int
	dictionary     string
	dnsServer      string
	dnsRecordTypes string
	domain         string
	log            logger.Configuration
	maxWordCount   int
	minWordSize    int
	maxWordSize    int
	output         string
	pattern        string
	source         string
	timeout        int
}
