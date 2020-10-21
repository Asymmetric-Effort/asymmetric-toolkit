package misc_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"testing"
)

func TestIntLen(t *testing.T) {
	errors.Assert(misc.IntLen(1) == 1, "Failed at 1")
	errors.Assert(misc.IntLen(10) == 2, "Failed at 10")
	errors.Assert(misc.IntLen(100) == 3, "Failed at 100")
	errors.Assert(misc.IntLen(1000) == 4, "Failed at 1000")
	errors.Assert(misc.IntLen(10000) == 5, "Failed at 10000")
	errors.Assert(misc.IntLen(100000) == 6, "Failed at 100000")
	errors.Assert(misc.IntLen(1000000) == 7, "Failed at 1000000")
}
