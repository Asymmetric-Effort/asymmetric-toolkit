package io

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDictionary_CompressionType(t *testing.T) {
	var o Compression
	errors.Assert(o == NoCompress, "expect 0")
	errors.Assert(NoCompress == 0, "Expect 0")
	errors.Assert(Gzip == 1, "expect 1")
}
