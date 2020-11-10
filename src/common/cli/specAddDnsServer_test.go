package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestSpecification_AddDnsServer(t *testing.T) {

	const testDnsServerDefault = "127.0.0.1"

	func() {
		var o Specification

		if dnsServerArgLong != "dnsServer" {
			panic("timeoutArgLong mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddDnsServer(testDnsServerDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[dnsServerArgLong].FlagId == FlagDnsServer,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d",
				dnsServerArgLong, FlagDnsServer, o.Argument[dnsServerArgLong].FlagId))

		errors.Assert(o.Argument[dnsServerArgLong].Type == String,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[dnsServerArgLong].Type))

		errors.Assert(o.Argument[dnsServerArgLong].Default == testDnsServerDefault,
			fmt.Sprintf("Default should be empty for %s", dnsServerArgLong))

		errors.Assert(o.Argument[dnsServerArgLong].Help == dnsServerHelpText,
			"Specifies the DNS server to be used by attackers.")

		errors.Assert(o.Argument[dnsServerArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[dnsServerArgLong].Expects == ExpectNone, "Next expected should be ExpectNone")

		if o.Argument[dnsServerArgLong].Default != testDnsServerDefault {
			panic(fmt.Sprintf("Default should be false for dnsServer flag (%d)", FlagDnsServer))
		}
	}()
}
