package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"fmt"
	"os"
	"testing"
)

func TestCreateDictionary_CliArgs(t *testing.T) {
	knownExistingFile := os.Args[0]
	tests := []struct {
		args    []string
		exit    int
		testMsg bool
		msg     string
	}{
		{ //test 0
			[]string{},
			cli.ErrArgumentParseError,
			true,
			"Error: " + cli.ErrMsgMissingArguments,
		}, { //test 1
			[]string{"--help"},
			cli.ErrSuccess,
			false,
			"",
		}, { //test 2
			[]string{"--help", "--debug"},
			cli.ErrSuccess,
			false,
			"",
		}, { //test 3
			[]string{"--version"},
			cli.ErrSuccess,
			false,
			"",
		}, { //test 4
			[]string{"--version", "--debug"},
			cli.ErrSuccess,
			false,
			"",
		}, { //test 5
			[]string{"--version", "--force"},
			cli.ErrSuccess,
			false,
			"",
		}, { //test 6
			[]string{"--in", ""},
			cli.ErrBadFilename,
			true,
			"Error: " + cli.ErrMsgEmptyInputFilename,
		}, { //test 7
			[]string{"--in", "nonexistent_input_file.txt"},
			cli.ErrFileNotFound,
			true,
			"Error: " + cli.ErrMsgInputFileNotFound,
		}, { //test 8
			[]string{"--in", knownExistingFile, "--out", ""},
			cli.ErrBadFilename,
			true,
			"Error: " + cli.ErrMsgEmptyOutputFilename,
		}, { //test 9
			[]string{"--in", knownExistingFile, "--out", knownExistingFile},
			cli.ErrBadFilename,
			true,
			"Error: " + cli.ErrMsgOutputFileExists,
		},
	}
	for i, test := range tests {
		exit, msg := createDictionary(test.args)
		if exit != test.exit {
			t.Errorf("Test %d: Expected exit %d but encountered %d", i, test.exit, exit)
			t.FailNow()
		}
		if test.testMsg && (msg != test.msg) {
			t.Errorf("Test %d: Expected msg '%s' but encountered '%s'", i, test.msg, msg)
			t.FailNow()
		}
		fmt.Printf("Test %d: OK [pass]\n", i)
	}
}

func TestCreateDictionary_TestFile(t *testing.T) {
	inputFile := "testInputFile.tmp"
	if file.FileExists(inputFile) {
		err := os.Remove(inputFile)
		if err != nil {
			panic(err)
		}
	}
	outputFile := "testOutputFile.tmp"
	if file.FileExists(outputFile) {
		err := os.Remove(outputFile)
		if err != nil {
			panic(err)
		}
	}
	defer func() {
		fmt.Println("Remove inputFile")
		err := os.Remove(inputFile)
		if err != nil {
			panic(err)
		}
	}()
	defer func() {
		fmt.Println("Remove outputFile")
		err := os.Remove(outputFile)
		if err != nil {
			panic(err)
		}
	}()
	// Create inputFile
	fmt.Println("Create Input file")
	f, err := os.Create(inputFile)
	if err != nil {
		panic(err)
	}
	defer closeFile(f)
	dataSet := []string{"word1", "word2", "word3", "word4", "word5", "word6", "word7", "word8"}
	for _, data := range dataSet {
		_, err = f.WriteString(data + "\n")
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("...input file created")
	//
	// Create output file.
	//
	fmt.Println("create output file")
	args := []string{"--in", inputFile, "--out", outputFile}
	exit, msg := createDictionary(args)
	if exit != cli.ErrSuccess {
		t.Errorf("createDirectory(): Expected ErrSuccess. inputFile:%s, outputFile:%s, msg:%s",
			inputFile, outputFile, msg)
		t.FailNow()
	}
	fmt.Println("...output file created.")
}
