package dictionary

import (
	"fmt"
)

func (o *ioCompression) Pack(in *[]byte) (out *[]byte) {
	switch *o {
	case noCompress:
		out = in

	case gzip:
		out = o.Gzip(in)
	default:
		panic(fmt.Sprintf("Unsupported ioCompression type (%d)", *o))
	}
	return
}
