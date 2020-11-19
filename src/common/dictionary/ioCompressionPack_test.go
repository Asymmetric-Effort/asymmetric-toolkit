package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestIoCompression_PackNoCompress(t *testing.T) {
	var o ioCompression = noCompress
	in := []byte("test")
	out := o.Pack(&in)
	errors.Assert(string(in) == string(*out), "noCompress expects input == output.")
}

func TestIoCompression_PackGzip(t *testing.T) {
	var o ioCompression = gzip
	in := []byte("test")
	out := o.Pack(&in)
	errors.Assert(string(in) != string(*out), "pack expects input != output.")
	out = o.Unpack(out)
	errors.Assert(string(in) == string(*out), "pack expects input == output.")
}
