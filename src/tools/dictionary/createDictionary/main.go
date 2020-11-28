package main

/*
	createDictionary is a tool for creating a dictionary with the asymmetric-toolkit dictionary format.
*/

import (
	"fmt"
	"os"
)

func closeFile(fileHandle *os.File) {
	err := fileHandle.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	exit, msg := createDictionary(os.Args[1:])
	fmt.Println(msg)
	os.Exit(exit)
}
