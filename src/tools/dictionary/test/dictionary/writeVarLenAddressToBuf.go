package dictionary

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func writeVarLenAddressToBuf(buf *bytes.Buffer, addr int64, sz uint32) (err error) {
	switch sz {
	case 1: //byte
		err = binary.Write(buf, binary.BigEndian, uint8(addr))
	case 2: //word16
		err = binary.Write(buf, binary.BigEndian, uint16(addr))
	case 4: //word32
		err = binary.Write(buf, binary.BigEndian, uint32(addr))
	case 8: //word64
		err = binary.Write(buf, binary.BigEndian, uint64(addr))
	default:
		err = fmt.Errorf("unsupported address size:%x", sz)
	}
	return
}
