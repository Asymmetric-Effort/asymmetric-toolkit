package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestProgramName(t *testing.T){
	errors.Assert(ProgramName=="dnsName","dnsName program name mismatch")
}