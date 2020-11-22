package io
/*

 */
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestIoCompression_GunZip(t *testing.T){
	var o Compression
	in := []byte("test")
	out := o.Gzip(&in)
	errors.Assert(string(in) != string(*out), "gzip expects input != output.")
	out = o.GunZip(out)
	errors.Assert(string(in) == string(*out), "gzip expects input == output.")
}