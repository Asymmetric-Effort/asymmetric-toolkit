package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strconv"
	"testing"
)

func TestSpecification_AddConcurrency(t *testing.T) {
	const defaultConcurrency = 10
	var o Specification

	if concurrencyArgLong != "concurrency" {
		panic("concurrencyArgLong mismatch")
	}

	errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

	o.AddConcurrency(defaultConcurrency)

	errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

	errors.Assert(o.Argument[concurrencyArgLong].FlagId == FlagConcurrency,
		fmt.Sprintf("(%s) expected.  FlagId:%d", concurrencyArgLong, o.Argument[concurrencyArgLong].FlagId))

	errors.Assert(o.Argument[concurrencyArgLong].Type == Integer,
		fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[concurrencyArgLong].Type))

	errors.Assert(o.Argument[concurrencyArgLong].Default == strconv.Itoa(defaultConcurrency),
		fmt.Sprintf("Default should be empty for %s", concurrencyArgLong))

	errors.Assert(o.Argument[concurrencyArgLong].Help == concurrencyHelpText, "usageHelpText mismatch")

	errors.Assert(o.Argument[concurrencyArgLong].Parse != nil, "Expect non-nil function pointer")

	errors.Assert(o.Argument[concurrencyArgLong].Expects == ExpectNone, "Next expected should be ExpectNone")

	val, err := strconv.Atoi(o.Argument[concurrencyArgLong].Default)
	if err != nil {
		panic(err)
	}
	if val != defaultConcurrency {
		panic(fmt.Sprintf("Default should be false for concurrency flag (%d): '%d'", defaultConcurrency, val))
	}
}
