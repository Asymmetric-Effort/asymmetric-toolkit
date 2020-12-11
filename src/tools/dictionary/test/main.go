package main

/*
	This is a test program to create a prototype of our dictionary file.
*/

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dictionary/test/dictionary"
	"fmt"
	"os"
)

const treeRootPtr = 0x12 //0x12

func main() {
	var sourceFile SourceFile
	var targetFile dictionary.File

	if file.FileExists(outputFile) {
		if err := os.Remove(outputFile); err != nil {
			panic(err)
		}
	}

	sourceFile.Open(inputFile)
	defer sourceFile.Close()

	targetFile.Open(outputFile)
	defer targetFile.Close()

	fmt.Println("Initial state:")
	fmt.Println("\tDictionary Header:   ", *targetFile.Header)
	fmt.Println("\tDictionary err:      ", targetFile.Err)
	fmt.Println("\tDictionary handle:   ", targetFile.Handle)
	fmt.Println("")

	fmt.Println("main(): Phase I - Import words (STARTING)")
	count := 0


	prevWordPtr = treeRootPtr
	for hasNext, word := sourceFile.Next(); hasNext; hasNext, word = sourceFile.Next() {
		fmt.Println("main(): fetching...", "hasNext:", hasNext, " word: ", word)
		targetFile.Add(&word, prevWordPtr)
		fmt.Println("main(): Word added without issue.")
		fmt.Println("")
		count++
	}
	fmt.Println("main(): ", count)
	errors.Assert(count == 34, "Expected 34 records to be created.")
	fmt.Println("main(): Phase I - Import words (COMPLETE)")

	fmt.Println("main(): Phase II - Read Tree and List (STARTING)")
	targetFile.DumpWords()
	fmt.Println("main(): Phase II - Read Tree and List (COMPLETE)")
}
