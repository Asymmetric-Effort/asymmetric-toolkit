package io
/*

 */
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestIoCompression_PackNoCompress(t *testing.T) {
	var o Compression = NoCompress
	in := []byte("test")
	out := o.Pack(&in)
	errors.Assert(string(in) == string(*out), "noCompress expects input == output.")
}

func TestIoCompression_PackGzip(t *testing.T) {
	var o Compression = Gzip
	in := []byte("test")
	out := o.Pack(&in)
	errors.Assert(string(in) != string(*out), "pack expects input != output.")
	out = o.Unpack(out)
	errors.Assert(string(in) == string(*out), "pack expects input == output.")
}
