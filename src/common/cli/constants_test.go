package cli_test

import (
	cli2 "asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDnsEnumCliConstants(t *testing.T){
	errors.Assert(cli2.Version == "0.0.1", "Expect version")
	errors.Assert(cli2.DefaultConcurrency == 1, "Expect 1")
	errors.Assert(cli2.DefaultDepth == 1, "Expect 1")
	errors.Assert(cli2.MaxDepth == 20, "expect 20")
	errors.Assert(cli2.DefaultTimeout == 60, "expect 60")
	errors.Assert(cli2.DefaultFilterPattern != ``, "Expect non-empty string")
	errors.Assert(cli2.DefaultDNSRecordTypes != "", "Expect non-empty string")
	errors.Assert(cli2.DNSChars != "", "expect non-empty string")
	errors.Assert(cli2.SourceBufferSz > 0, "expect int > 0")
	errors.Assert(cli2.ResponseBufferSz > 0, "expect int > 0")
	errors.Assert(cli2.ExpectFlag == 0, "expect 0")
	errors.Assert(cli2.ExpectValue == 1, "expect 1")
	errors.Assert(cli2.ExitTerminate, "expect true")
	errors.Assert(!cli2.ExitParseOk, "expect false")
}