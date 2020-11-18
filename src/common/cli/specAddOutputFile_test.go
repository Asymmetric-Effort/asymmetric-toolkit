package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestCommandLine_AddOutputFile(t *testing.T) {
	var o Specification
	o.AddOutputFile()
	errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor")
	errors.Assert(o.Argument[outFileArgLong].FlagId == FlagOutputFile,
		fmt.Sprintf("(%s) expected (%d)  FlagId:%d",
			outFileArgLong, FlagOutputFile, o.Argument[outFileArgLong].FlagId))
	errors.Assert(o.Argument[outFileArgLong].Type == String, "Expected string type")
	errors.Assert(o.Argument[outFileArgLong].Expects == ExpectValue, "Expects a value")
	errors.Assert(o.Argument[outFileArgLong].Default == "", "Expect empty string")
	errors.Assert(o.Argument[outFileArgLong].Help == outFileHelpText, "expect help text.")
	errors.Assert(o.Argument[outFileArgLong].Parse != nil, "expect non nil parse function.")
}
