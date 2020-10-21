package destination_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/destination"
	"testing"
)

func TestLogDestinationGetHappy(t *testing.T){
	var ld destination.LogDestination
	ld= destination.Stdout
	errors.Assert(ld.Get() == destination.Stdout, "Expected Stdout")
	ld= destination.File
	errors.Assert(ld.Get() == destination.File, "Expected File")
	ld= destination.Syslog
	errors.Assert(ld.Get() == destination.Syslog, "Expected Syslog")
}
func TestLogDestinationGetSad(t *testing.T){
	var ld destination.LogDestination = 9
	defer func(){recover()}()
	_ = ld.Get()
	t.Fatal("expected error")
}