package utils

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestInt64Len(t *testing.T) {
	errors.Assert(1 == IntLen64(1), "Failed 1")
	errors.Assert(2 == IntLen64(10), "Failed 10")
	errors.Assert(3 == IntLen64(100), "Failed 100")
	errors.Assert(4 == IntLen64(1000), "Failed 1000")
	errors.Assert(5 == IntLen64(10000), "Failed 10000")
	errors.Assert(6 == IntLen64(100000), "Failed 100000")
	errors.Assert(7 == IntLen64(1000000), "Failed 1000000")
}
