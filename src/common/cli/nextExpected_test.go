package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestNextExpected(t *testing.T) {
	errors.Assert(ExpectFlag == 0, "Expected 0")
	errors.Assert(ExpectValue == 1, "Expected 1")
	errors.Assert(ExpectEnd == 2, "Expected 2")
	errors.Assert(ExpectNone == 3, "Expected 3")
}
