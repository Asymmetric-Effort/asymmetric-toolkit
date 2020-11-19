package main

/*
	createDictionary is a tool for creating a dictionary with the asymmetric-toolkit dictionary format.
*/

import (
	"fmt"
	"os"
)

func main() {
	exit, msg := createDictionary(os.Args[1:])
	fmt.Println(msg)
	os.Exit(exit)
}
