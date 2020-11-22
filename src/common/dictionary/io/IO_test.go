package io

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestIO(t *testing.T) {
	var i IO
	errors.Assert(i.Compress == NoCompress, "expect 0")
}
