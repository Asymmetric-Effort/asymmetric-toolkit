package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestArgument(t *testing.T){
	var o Argument
	errors.Assert(o.Type == None, "Expected zero ArgumentType (None)")
	errors.Assert(o.Value == "", "expected empty string")
}