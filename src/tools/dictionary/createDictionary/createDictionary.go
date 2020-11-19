package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"bufio"
	"fmt"
	"os"
)



type DefinitionScore int

type SourceFile struct {
	name      string
	handle    *os.File
	hashtable HashTable
}

type DictionaryHeader struct {
	version uint32
}

type DictionaryDefinition struct {
	Id      []byte
	Word    string
	Score   DefinitionScore
	Origin  DefinitionOrigin
	Type    DefinitionType
	Created int64
	LastHit int64
}

type DictionaryFile struct {
	name    string
	handle  *os.File
	header  DictionaryHeader
	content []DictionaryDefinitions
	footer  DictionaryFooter
}

func closeFile(fileHandle *os.File) {
	err := fileHandle.Close()
	if err != nil {
		panic(err)
	}
}
func writeDefinition(fileHandle *os.File, config *Configuration) (exit int, msg string) {
	//
	return
}

func writeHeader(fileHandle *os.File, config *Configuration) (exit int, msg string) {
	//
	return
}

func writeFooter(fileHandle *os.File, config *Configuration) (exit int, msg string) {
	//
	return
}

func openOutputFile(fileName string) *os.File {
	f, err := os.Create(config.OutputFile)
	if err != nil {
		panic(err)
	}
	return f
}

func createDictionary(args []string) (exit int, msg string) {
	exit = cli.ErrSuccess // Return with no errors
	msg = "OK"
	/*
		Parse the commandline arguments and configure our internal state.
	*/
	config, exitProgram, err := ProcessSpecification(args)
	switch {
	/*
		Verify the commandline arguments are okay.
	*/
	case err != nil:
		return cli.ErrArgumentParseError, fmt.Sprintf("Error: %v", err)
	case exitProgram || (config == nil):
		return cli.ErrSuccess, "Exit Success"
	}
	/*
		Configure our logger.
	*/
	//goland:noinspection GoNilness
	var log logger.Logger
	log.Setup(&config.Log)
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

	sourceFile, err := os.Open(config.InputFile)
	if err != nil {
		return cli.ErrFileOpenFailed, fmt.Sprintf("Failed to Open file (Error: %v)", err)
	}
	defer closeFile(sourceFile)

	outputFile := openOutputFile(config.OutputFile)
	defer closeFile(outputFile)

	exit, msg = writeHeader(&outputFile, config)
	if exit != cli.ErrSuccess {
		return exit, msg
	}

	var wordHash map[hash]struct{}
	{
	}

	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return cli.ErrFileRead, fmt.Sprintf("failed to read from file (Error: %v)", err)
		}
		word := scanner.Text()
		exit, msg = writeDefinition(&outputFile, word)
		if exit != cli.ErrSuccess {
			return exit, msg
		}
	}
	exit, msg = writeFooter(&outputFile, config)
	return
}
