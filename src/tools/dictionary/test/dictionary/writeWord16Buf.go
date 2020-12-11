package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bytes"
	"encoding/binary"
	"fmt"
)

func writeWord16Buf(buf *bytes.Buffer, data int) (err error) {
	//
	// Write string size (up to 65535 length) to buffer.
	//
	errors.Assert(data < 65536, "Expect word size < 65536.")
	if err = binary.Write(buf, binary.BigEndian, uint16(data)); err != nil {
		return fmt.Errorf("Error  writing WordSz (%x):%v", data, err)
	}
	fmt.Println("\tFile::WriteWord(): word size written to buffer")
	return nil
}
