package cli

import (
	"fmt"
	"testing"
)

func TestSpecification_AddDebug(t *testing.T) {

	var o Specification

	if debugArgLong != "debug" {
		panic("debugArgLong mismatch")
	}

	if o.Argument != nil {
		panic("Expected nil ArgumentDescriptor in Specification.")
	}

	o.AddDebug()

	if o.Argument == nil {
		panic("Expected nil ArgumentDescriptor in Specification.")
	}

	if o.Argument[debugArgLong].FlagId != flagDebug {
		panic(fmt.Sprintf("(%s) expected.  FlagId:%d", debugArgLong, o.Argument[debugArgLong].FlagId))
	}

	if o.Argument[debugArgLong].Type != Boolean {
		panic(fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[debugArgLong].Type))
	}

	if o.Argument[debugArgLong].Default != debugDefault {
		panic(fmt.Sprintf("Default should be empty for %s", debugArgLong))
	}

	if o.Argument[debugArgLong].Help != debugHelpText {
		panic("usageHelpText mismatch")
	}

	if o.Argument[debugArgLong].Parse == nil {
		panic("Expect non-nil function pointer")
	}

	if o.Argument[debugArgLong].Expects != ExpectNone {
		panic("Next expected should be ExpectNone")
	}
}
