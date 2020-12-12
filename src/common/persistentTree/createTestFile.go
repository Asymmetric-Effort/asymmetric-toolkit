package persistentTree

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"encoding/binary"
	"os"
)

func createTestFile(n string) {
	/*
		create a test file with default header data.
	*/
	h, err := os.Create(n)
	if err != nil {
		panic(err)
	}
	defer file.CloseFile(h)

	err = binary.Write(h, binary.BigEndian, fileHeaderVersion)
	if err != nil {
		panic(err)
	}

	err = binary.Write(h, binary.BigEndian, uint16(0)) // Flags.
	if err != nil {
		panic(err)
	}

	err = binary.Write(h, binary.BigEndian, uint16(12)) // Flags.
	if err != nil {
		panic(err)
	}
}