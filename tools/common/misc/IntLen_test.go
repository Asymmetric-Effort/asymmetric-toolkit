package misc

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestIntLen(t *testing.T) {
	errors.Assert(1 == IntLen(1), "Failed 1")
	errors.Assert(2 == IntLen(10), "Failed 10")
	errors.Assert(3 == IntLen(100), "Failed 100")
	errors.Assert(4 == IntLen(1000), "Failed 1000")
	errors.Assert(5 == IntLen(10000), "Failed 10000")
	errors.Assert(6 == IntLen(100000), "Failed 100000")
	errors.Assert(7 == IntLen(1000000), "Failed 1000000")
}
