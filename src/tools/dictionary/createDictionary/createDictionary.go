package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary"
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"bufio"
	"fmt"
	"os"
)

type Dictionary struct {
	handle   *os.File
	position uint32

	header DictionaryFileHeader
}

type DictionaryFileHeader struct {
	FileVersion  uint16 //
	ScoreVersion uint16 //
	ReservedA    uint8  //                        +--------------------------=+
	ReservedB    uint8  //                        |  Dictionary Header Flags  |
	ReservedC    uint8  //                        +---------------------------+
	//						//           F P          |FOE - File Operation Error |
	//						//           O F F        |PFL - Pending File Lock    |
	//						// ----------E L L        | FL - File Lock            |
	Flags            uint8  // 7 6 5 4 3 2 1 0        +--------------------------=+
	TagTableOffset   uint32 // Offset for TagTable
	DefinitionOffset uint32 // Offset for DefinitionTable
}

func createDictionary(args []string) (exit int, msg string) {
	var log logger.Logger
	exit, msg = cli.ErrSuccess, "OK" // Return with no errors
	//
	//	Parse the commandline arguments and configure our internal state.
	//
	config, exitProgram, err := ProcessSpecification(args)
	switch {
	//
	//	Verify the commandline arguments are okay.
	//
	case err != nil:
		return cli.ErrArgumentParseError, fmt.Sprintf("Error: %v", err)
	case exitProgram || (config == nil):
		return cli.ErrSuccess, "Exit Success"
	}
	//
	// Evaluate inputs
	//
	switch {
	case config.InputFile == "":
		return cli.ErrBadFilename, "Error: " + cli.ErrMsgEmptyInputFilename

	case !file.FileExists(config.InputFile):
		return cli.ErrFileNotFound, "Error: " + cli.ErrMsgInputFileNotFound

	case config.OutputFile == "":
		return cli.ErrBadFilename, "Error: " + cli.ErrMsgEmptyOutputFilename

	case file.FileExists(config.OutputFile) && !config.Force:
		return cli.ErrBadFilename, "Error: " + cli.ErrMsgOutputFileExists
	}
	//
	// Configure our logger.
	//
	log.Setup(&config.Log)
	//
	// Open the source file.
	//
	sourceFile, err := os.Open(config.InputFile)
	if err != nil {
		return cli.ErrFileOpenFailed, fmt.Sprintf("Failed to Open file (Error: %v)", err)
	}
	defer closeFile(sourceFile)

	var dictionary dictionary.File

	err := dictionary.Setup(config)
	if err != nil {
		panic(err)
	}
	defer dictionary.Teardown()

	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return cli.ErrFileRead, fmt.Sprintf("failed to read from file (Error: %v)", err)
		}
		word := scanner.Text()
		dictionary.WriteWord(word)
	}
	return
}
