package headers

import (
	"encoding/binary"
)

func (o *Descriptor) SetHeader(header *Descriptor) {

	err := binary.Write(o.FileHandle, binary.LittleEndian, &o.Header)

	if err != nil {
		panic(err)
	}

}