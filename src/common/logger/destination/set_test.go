package destination_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/destination"
	"testing"
)

func TestLogDestinationSetHappy(t *testing.T){
	var ld destination.LogDestination
	ld.Set(destination.Stdout)
	errors.Assert(ld.Get() == destination.Stdout, "Expected Stdout")
	ld.Set(destination.File)
	errors.Assert(ld.Get() == destination.File, "Expected File")
	ld.Set(destination.Syslog)
	errors.Assert(ld.Get() == destination.Syslog, "Expected Syslog")
}
func TestLogDestinationSetSad(t *testing.T){
	var ld destination.LogDestination
	defer func(){recover()}()
	ld.Set(9)
	t.Fatal("expected error")
}