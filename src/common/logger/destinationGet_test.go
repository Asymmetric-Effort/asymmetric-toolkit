package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogDestinationGetHappy(t *testing.T){
	var ld Destination
	ld= Stdout
	errors.Assert(ld.Get() == Stdout, "Expected Stdout")
	ld= File
	errors.Assert(ld.Get() == File, "Expected File")
	ld= Syslog
	errors.Assert(ld.Get() == Syslog, "Expected Syslog")
}

func TestLogDestinationGetSad(t *testing.T){
	var ld Destination = 9
	defer func(){recover()}()
	_ = ld.Get()
	t.Fatal("expected error")
}
