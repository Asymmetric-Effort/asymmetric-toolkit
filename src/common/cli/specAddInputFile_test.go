package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestCommandline_AddInputFile(t *testing.T){
	var o Specification
	o.AddInputFile()
	errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor")
	errors.Assert(o.Argument[inFileArgLong].FlagId == FlagInputFile,
		fmt.Sprintf("(%s) expected (%d)  FlagId:%d",
			inFileArgLong, FlagInputFile, o.Argument[inFileArgLong].FlagId))
	errors.Assert(o.Argument[inFileArgLong].Type == String, "Expected string type")
	errors.Assert(o.Argument[inFileArgLong].Expects == ExpectValue, "Expects a value")
	errors.Assert(o.Argument[inFileArgLong].Default == "", "Expect empty string")
	errors.Assert(o.Argument[inFileArgLong].Help == inFileHelpText, "expect help text.")
	errors.Assert(o.Argument[inFileArgLong].Parse != nil, "expect non nil parse function.")
}
