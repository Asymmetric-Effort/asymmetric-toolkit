package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestCliConstants(t *testing.T){
	errors.Assert(ErrSuccess == 0, "Expected 0")
	errors.Assert(ErrArgumentParseError == 1, "Expected 1")
}