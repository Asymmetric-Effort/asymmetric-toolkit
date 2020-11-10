package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestArgumentDescriptor(t *testing.T) {
	var o ArgumentDescriptor
	errors.Assert(o.FlagId == noFlag, "Expect noFlag")
	errors.Assert(o.Type == None, "Expect None ArgumentType")
	errors.Assert(o.Default == "", "default to empty string")
	errors.Assert(o.Help == "", "Expect empty string")
	errors.Assert(o.Parse == nil, "Expect nil pointer")
	errors.Assert(o.Expects == ExpectFlag, "Expect flag")
}
