package misc

import "encoding/binary"

const (
	uint32Sz = 4 //bytes
)

func Uint32ToByte(i uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	return b
}
