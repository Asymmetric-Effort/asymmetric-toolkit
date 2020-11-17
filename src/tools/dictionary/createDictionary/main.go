package main

/*
	createDictionary is a tool for creating a dictionary with the asymmetric-toolkit dictionary format.
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
	log.Setup(&config.Log)
	log.Debug(logger.EventInit)

	// the next line will just exit the application.
	exit <- cli.ErrSuccess
	//
	// Payload.
	//	Open the input file.
	//	Open the output file (if --force, overwrite the file, otherwise stop if it exists).
	//	Create the file header struct.
	//	For each line in the input file...
	//		Create a dictionary definition and write it to a temporary file.
	//	Add the header to the output file and then copy the contents of the temporary file into the output file.
	//

	<-exit
}
