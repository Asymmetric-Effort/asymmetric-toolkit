package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestTagType(t *testing.T) {
	errors.Assert(BasicTag == 0, "expect 0")
	errors.Assert(BooleanTag == 1, "expect 1")
	errors.Assert(StringTag == 2, "expect 2")
	errors.Assert(IntegerTag == 3, "expect 3")
	errors.Assert(Float64Tag == 4, "expect 4")
}
