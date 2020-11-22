package io
/*

 */
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestIoCompression_Gzip(t *testing.T){
	var o Compression
	in := []byte("test")
	fmt.Println("Stage 1 (start)")
	out := o.Gzip(&in)
	fmt.Println("Stage 1 (end)")
	errors.Assert(string(in) != string(*out), "gzip expects input != output.")
	fmt.Println("Stage 2 (start)")
	out = o.GunZip(out)
	fmt.Println("Stage 2 (end)")
	errors.Assert(string(in) == string(*out), "gzip expects input == output.")
}