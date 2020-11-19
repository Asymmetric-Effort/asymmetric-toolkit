package dictionary

import "encoding/binary"

func (o *Dictionary) SetHeader(header *HeaderDescriptor) {
	err := binary.Write(o.fileHandle, binary.LittleEndian, &o.Content.Header)
	if err != nil {
		panic(err)
	}
}