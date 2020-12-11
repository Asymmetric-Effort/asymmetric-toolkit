package file

import "os"

func CloseFile(fileHandle *os.File) {
	if fileHandle == nil {
		return
	}
	err := fileHandle.Close()
	if err != nil {
		panic(err)
	}
}