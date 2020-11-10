package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestConstants(t *testing.T) {

	errors.Assert(ProgramName == "dnsName", "ProgramName mismatch")

	errors.Assert(DnsRecordTypes == "A,AAAA,MX,TXT,SOA,NS,CNAME", "DnsRecordTypes mismatch")

	errors.Assert(DefaultConcurrency == 1, "DNS Concurrency mismatch")

	errors.Assert(DefaultAttackDepth == 20, "DNS Depth mismatch")

	errors.Assert(DefaultTimeout == 60, "DefaultTimeout mismatch")

	errors.Assert(DefaultFilterPattern == ".+", "DefaultFilterPattern mismatch")

	errors.Assert(DefaultMinWordSize == 3, "DefaultMinWordSize mismatch")

	errors.Assert(DefaultMaxWordSize == 5, "DefaultMaxWordSize mismatch")

	//errors.Assert(SourceBufferSz == 1048676, "default source buffer size.")

	//errors.Assert(ResponseBufferSz == 1048576, "default response buffer size.")

	errors.Assert(DnsChars == "WMEFSABCDGHIJKLNOPQRTUVXYZwmefsabcdghijklnopqrtuvxyz0123456789._-",
		"dns character mismatch")

}
