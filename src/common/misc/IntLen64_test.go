package misc_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"testing"
)

func TestInt64Len(t *testing.T) {
	errors.Assert(1 == misc.IntLen64(1), "Failed 1")
	errors.Assert(2 == misc.IntLen64(10), "Failed 10")
	errors.Assert(3 == misc.IntLen64(100), "Failed 100")
	errors.Assert(4 == misc.IntLen64(1000), "Failed 1000")
	errors.Assert(5 == misc.IntLen64(10000), "Failed 10000")
	errors.Assert(6 == misc.IntLen64(100000), "Failed 100000")
	errors.Assert(7 == misc.IntLen64(1000000), "Failed 1000000")
}
