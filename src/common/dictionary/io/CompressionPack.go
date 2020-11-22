package io
/*
	Compression::Pack() is the top-level compression method.  Any compression calls should pass through to this
                        method, and depending on the state of the Compression object, execution flows to the
                        specific compression algorithm method.
 */
import (
	"fmt"
)

func (o *Compression) Pack(in *[]byte) (out *[]byte) {
	switch *o {
	case NoCompress:
		out = in

	case Gzip:
		out = o.Gzip(in)
	default:
		panic(fmt.Sprintf("Unsupported Compression type (%d)", *o))
	}
	return
}
