package dictionary

import (
	"fmt"
)

func (o *ioCompression) Pack(in *[]byte) (out *[]byte) {
	switch *o {
	case NoCompress:
		out = in

	case Gzip:
		out = o.Gzip(in)
	default:
		panic(fmt.Sprintf("Unsupported ioCompression type (%d)", *o))
	}
	return
}
