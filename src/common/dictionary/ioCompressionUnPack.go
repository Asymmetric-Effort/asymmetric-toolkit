package dictionary

import (
	"fmt"
)

func(o *ioCompression)Unpack(in *[]byte)(out *[]byte){
	switch *o {
	case noCompress:
		out = in
	case gzip:
		out = o.GunZip(in)
	default:
		panic(fmt.Sprintf("Unsupported ioCompression type (%d)",*o))
	}
	return
}
