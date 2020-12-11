package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"bufio"
	"fmt"
	"os"
)

type SourceFile struct {
	handle  *os.File
	err     error
	scanner *bufio.Scanner
}

func (o *SourceFile) Open(filename string) {
	o.handle, o.err = os.Open(filename)
	errors.Assert(o.err == nil, fmt.Sprintf("%v", o.err))
	o.scanner = bufio.NewScanner(o.handle)
	return
}
func (o *SourceFile) Close() {
	file.CloseFile(o.handle)
}

func (o *SourceFile) Next() (hasWord bool, word string) {
	if o.scanner.Scan() {
		o.err = o.scanner.Err()
		errors.Assert(o.err == nil, fmt.Sprintf("%v", o.err))
		word = o.scanner.Text()
		hasWord = true
	}
	fmt.Println("SourceFile::Next() ", hasWord, " : ", word)
	return
}
