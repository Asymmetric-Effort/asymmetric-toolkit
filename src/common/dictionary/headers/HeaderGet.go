package headers

import (
	"encoding/binary"
)

func (o *Descriptor) GetHeader() (header *Descriptor, err error) {
	err = binary.Read(o.fileHandle, binary.LittleEndian, &o.Content.Header)
	return
}
