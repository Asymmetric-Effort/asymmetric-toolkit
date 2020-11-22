package io
/*
	Compression::Unpack() is the top-level decompression method.  Any decompression operation should pass through this
                          method, and depending on the state of the compression object, the execution will flow to the
                          specific decompression method.
 */
import (
	"fmt"
)

func(o *Compression)Unpack(in *[]byte)(out *[]byte){
	switch *o {
	case NoCompress:
		out = in
	case Gzip:
		out = o.GunZip(in)
	default:
		panic(fmt.Sprintf("Unsupported Compression type (%d)",*o))
	}
	return
}
