package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDnsEnumCliConstants(t *testing.T){
	errors.Assert(Version == "0.0.1", "Expect version")
	errors.Assert(defaultConcurrency == 1, "Expect 1")
	errors.Assert(defaultDepth == 1, "Expect 1")
	errors.Assert(maxDepth == 20, "expect 20")
	errors.Assert(defaultTimeout == 60, "expect 60")
	errors.Assert(defaultFilterPattern != ``, "Expect non-empty string")
	errors.Assert(defaultDnsRecordTypes != "", "Expect non-empty string")
	errors.Assert(DnsChars != "", "expect non-empty string")
	errors.Assert(SourceBufferSz > 0, "expect int > 0")
	errors.Assert(ResponseBufferSz > 0, "expect int > 0")
	errors.Assert(expectFlag == 0, "expect 0")
	errors.Assert(expectValue == 1, "expect 1")
	errors.Assert(ExitTerminate, "expect true")
	errors.Assert(!ExitParseOk, "expect false")
}