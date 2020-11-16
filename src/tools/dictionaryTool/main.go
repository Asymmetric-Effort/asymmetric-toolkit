package main

/*
	dictionaryTool is a tool for creating, maintaining and editing Asymmetric Toolkit dictionary files.
	These dictionary files are formatted with not only the encrypted word contained in the dictionary as well
	as metadata descriptive of the word (e.g. hits, discovery date, last date, etc.).
*/

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"os"
)

func main() {
	//
	// Setup our program to run
	//
	var exit chan int = make(chan int, 1) // We will block until an exit code is written to this channel.
	config, exitProgram, err := ProcessSpecification(os.Args[1:])
	if err != nil {
		exit <- cli.ErrArgumentParseError
	}
	if exitProgram {
		exit <- cli.ErrSuccess
	}
	errors.Assert(config == nil, "Internal error nil config encountered.")

	var log logger.Logger
	log.Setup(&config.log)
	log.Debug(logger.EventInit)

	// the next line will just exit the application.
	exit <- cli.ErrSuccess
	//
	// Payload.
	//

	<-exit
}
