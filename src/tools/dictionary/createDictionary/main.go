package main

/*
	createDictionary is a tool for creating a dictionary with the asymmetric-toolkit dictionary format.
*/

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"bufio"
	"fmt"
	"os"
)

func main() {
	//
	// Setup our program to run
	//
	var log logger.Logger
	var exit = make(chan int, 1) // We will block until an exit code is written to this channel.
	config, exitProgram, err := ProcessSpecification(os.Args[1:])
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(cli.ErrArgumentParseError)
	}
	if exitProgram || (config == nil) {
		os.Exit(cli.ErrSuccess)
	}
	//goland:noinspection GoNilness
	log.Setup(&config.Log)

	if config.InputFile == "" {
		fmt.Println("Error: empty input filename")
		os.Exit(cli.ErrBadFilename)
	}
	if !file.FileExists(config.InputFile) {
		fmt.Println("Error: input file not found")
		os.Exit(cli.ErrFileNotFound)
	}
	if config.OutputFile == "" {
		fmt.Println("Error: empty output filename")
		os.Exit(cli.ErrBadFilename)
	}
	if file.FileExists(config.OutputFile) && !config.Force {
		fmt.Println("Error: output file exists (--force required to overwrite)")
		os.Exit(cli.ErrBadFilename)
	}
	sourceFile, err := os.Open(config.InputFile)
	if err != nil {
		fmt.Printf("Failed to Open file (Error: %v)", err)
		os.Exit(cli.ErrFileOpenFailed)
	}
	defer func() {
		err := sourceFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	for scanner := bufio.NewScanner(sourceFile); scanner != nil; scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("Failed to read from file (Error: %v)", err)
			os.Exit(cli.ErrFileRead)
		}
		word:=scanner.Text()
		fmt.Println(word)
	}

	//
	// Open input file
	// Create file header
	// Read the file line by line
	// 		Create a definition structure and load it with the word data from the source file.
	// 		Score the definition
	//		Write the definition to the output file.
	// Close input file.
	// Create footer
	// Write footer to output file
	// Close output file.
	//

	exit <- cli.ErrSuccess
	os.Exit(<-exit)
}
