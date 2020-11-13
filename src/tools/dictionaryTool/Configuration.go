package main

import "asymmetric-effort/asymmetric-toolkit/src/common/logger"

/*
	Configuration is a final internal state of the application after the
	command line arguments are parsed and processed.
*/

type Configuration struct {
	dictionary     string
	log            logger.Configuration
}
