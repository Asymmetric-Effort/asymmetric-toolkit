package misc

import "encoding/binary"

const (
	uint64Sz = 8 //bytes
)

func Uint64ToByte(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}
