package main

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
	maxWordCount   int
	minWordSize    int
	maxWordSize    int
	output         string
	pattern        string
	source         string
	timeout        int
}
