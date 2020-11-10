package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestSpecification_AddDnsRecordType(t *testing.T) {

	func() {
		const TestDnsRecordTypeDefault = "A,CNAME"
		var o Specification

		if dnsRecordTypesArgLong != "dnsRecordTypes" {
			panic("dnsRecordType mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddDnsRecordType(TestDnsRecordTypeDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[dnsRecordTypesArgLong].FlagId == FlagDnsRecordType,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d", dnsRecordTypesArgLong, FlagDnsRecordType,
				o.Argument[dnsRecordTypesArgLong].FlagId))

		errors.Assert(o.Argument[dnsRecordTypesArgLong].Type == List,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[dnsRecordTypesArgLong].Type))

		errors.Assert(o.Argument[dnsRecordTypesArgLong].Default == TestDnsRecordTypeDefault,
			fmt.Sprintf("Default should be empty for %s", dnsRecordTypesArgLong))

		errors.Assert(o.Argument[dnsRecordTypesArgLong].Help == dnsRecordTypesHelpText,
			"Specifies a comma-delimited list of dns record types.")

		errors.Assert(o.Argument[dnsRecordTypesArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[dnsRecordTypesArgLong].Expects == ExpectValue, "Next expected should be ExpectValue")
	}()
}
