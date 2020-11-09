package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strconv"
	"testing"
)

func TestSpecification_AddDebug(t *testing.T) {

	var o Specification

	if debugArgLong != "debug" {
		panic("debugArgLong mismatch")
	}

	errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

	o.AddDebug()

	errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

	errors.Assert(o.Argument[debugArgLong].FlagId == flagDebug,
		fmt.Sprintf("(%s) expected.  FlagId:%d", debugArgLong, o.Argument[debugArgLong].FlagId))

	errors.Assert(o.Argument[debugArgLong].Type == Boolean,
		fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[debugArgLong].Type))

	errors.Assert(o.Argument[debugArgLong].Default == debugDefault,
		fmt.Sprintf("Default should be empty for %s", debugArgLong))

	errors.Assert(o.Argument[debugArgLong].Help == debugHelpText, "usageHelpText mismatch")

	errors.Assert(o.Argument[debugArgLong].Parse != nil, "Expect non-nil function pointer")

	errors.Assert(o.Argument[debugArgLong].Expects == ExpectNone, "Next expected should be ExpectNone")

	val, err := strconv.ParseBool(o.Argument[debugArgLong].Default)

	if err != nil {
		panic(err)
	}

	if val {
		panic("Default should be false for debug flag.")
	}
}
