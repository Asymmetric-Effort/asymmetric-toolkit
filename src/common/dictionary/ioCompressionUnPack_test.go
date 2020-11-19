package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestIoCompression_UnPackNoCompress(t *testing.T) {
	var o ioCompression = noCompress
	in := []byte("test")
	out := o.Pack(&in)
	errors.Assert(string(in) == string(*out), "noCompress expects input == output.")
}

func TestIoCompression_UnPackGzip(t *testing.T) {
	var o ioCompression = gzip
	in := []byte("test")
	out := o.Pack(&in)
	errors.Assert(string(in) != string(*out), "unpack expects input != output.")
	out = o.Unpack(out)
	errors.Assert(string(in) == string(*out), "unpack expects input == output.")
}
