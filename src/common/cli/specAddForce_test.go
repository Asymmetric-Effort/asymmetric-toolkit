package cli

import (
	"fmt"
	"testing"
)

func TestSpecification_AddForce(t *testing.T) {

	var o Specification

	if forceArgLong != "force" {
		panic("forceArgLong mismatch")
	}

	if o.Argument != nil {
		panic("Expected nil ArgumentDescriptor in Specification.")
	}

	o.AddForce()

	if o.Argument == nil {
		panic("Expected nil ArgumentDescriptor in Specification.")
	}

	if o.Argument[forceArgLong].FlagId != flagForce {
		panic(fmt.Sprintf("(%s) expected.  FlagId:%d", forceArgLong, o.Argument[forceArgLong].FlagId))
	}

	if o.Argument[forceArgLong].Type != Boolean {
		panic(fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[forceArgLong].Type))
	}

	if o.Argument[forceArgLong].Default != forceDefault {
		panic(fmt.Sprintf("Default should be empty for %s", forceArgLong))
	}

	if o.Argument[forceArgLong].Help != forceHelpText {
		panic("usageHelpText mismatch")
	}

	if o.Argument[forceArgLong].Parse == nil {
		panic("Expect non-nil function pointer")
	}

	if o.Argument[forceArgLong].Expects != ExpectNone {
		panic("Next expected should be ExpectNone")
	}
}
