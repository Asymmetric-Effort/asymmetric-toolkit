package dictionary

import (
	"fmt"
)

func(o *ioCompression)Unpack(in *[]byte)(out *[]byte){
	switch *o {
	case NoCompress:
		out = in
	case Gzip:
		out = o.GunZip(in)
	default:
		panic(fmt.Sprintf("Unsupported ioCompression type (%d)",*o))
	}
	return
}
