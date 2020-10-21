package cli_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
	"testing"
)

func TestDnsEnumCliConstants(t *testing.T){
	errors.Assert(cli.Version == "0.0.1", "Expect version")
	errors.Assert(cli.DefaultConcurrency == 1, "Expect 1")
	errors.Assert(cli.DefaultDepth == 1, "Expect 1")
	errors.Assert(cli.MaxDepth == 20, "expect 20")
	errors.Assert(cli.DefaultTimeout == 60, "expect 60")
	errors.Assert(cli.DefaultFilterPattern != ``, "Expect non-empty string")
	errors.Assert(cli.DefaultDNSRecordTypes != "", "Expect non-empty string")
	errors.Assert(cli.DNSChars != "", "expect non-empty string")
	errors.Assert(cli.SourceBufferSz > 0, "expect int > 0")
	errors.Assert(cli.ResponseBufferSz > 0, "expect int > 0")
	errors.Assert(cli.ExpectFlag == 0, "expect 0")
	errors.Assert(cli.ExpectValue == 1, "expect 1")
	errors.Assert(cli.ExitTerminate, "expect true")
	errors.Assert(!cli.ExitParseOk, "expect false")
}