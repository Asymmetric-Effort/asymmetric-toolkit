package dictionary

import "encoding/binary"

func (o *Dictionary) GetHeader() (header *HeaderDescriptor, err error) {
	err = binary.Read(o.fileHandle, binary.LittleEndian, &o.Content.Header)
	return
}
