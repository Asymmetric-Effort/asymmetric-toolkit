package misc_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"testing"
)

func TestIntLen(t *testing.T) {
	errors.Assert(1 == misc.IntLen(1), "Failed 1")
	errors.Assert(2 == misc.IntLen(10), "Failed 10")
	errors.Assert(3 == misc.IntLen(100), "Failed 100")
	errors.Assert(4 == misc.IntLen(1000), "Failed 1000")
	errors.Assert(5 == misc.IntLen(10000), "Failed 10000")
	errors.Assert(6 == misc.IntLen(100000), "Failed 100000")
	errors.Assert(7 == misc.IntLen(1000000), "Failed 1000000")
}
