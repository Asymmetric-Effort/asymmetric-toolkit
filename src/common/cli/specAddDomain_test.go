package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestSpecification_AddDomain(t *testing.T) {

	const testDefault = "localdomain"

	func() {
		var o Specification

		if domainArgLong != "domain" {
			panic("mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddDomain(testDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[domainArgLong].FlagId == FlagDomain,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d",
				domainArgLong, FlagDomain, o.Argument[domainArgLong].FlagId))

		errors.Assert(o.Argument[domainArgLong].Type == String,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[domainArgLong].Type))

		errors.Assert(o.Argument[domainArgLong].Default == testDefault,
			fmt.Sprintf("Default should be empty for %s", domainArgLong))

		errors.Assert(o.Argument[domainArgLong].Help == domainHelpText, "Specifies a fully qualified domain name.")

		errors.Assert(o.Argument[domainArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[domainArgLong].Expects == ExpectValue, "Next expected should be ExpectValue")

		if o.Argument[domainArgLong].Default != testDefault {
			panic(fmt.Sprintf("Default should be false for domain flag (%d)", FlagDomain))
		}
	}()
}
