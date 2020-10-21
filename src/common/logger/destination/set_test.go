package destination_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogDestinationSetHappy(t *testing.T){
	var ld LogDestination
	ld.Set(Stdout)
	errors.Assert(ld.Get() == Stdout, "Expected Stdout")
	ld.Set(File)
	errors.Assert(ld.Get() == File, "Expected File")
	ld.Set(Syslog)
	errors.Assert(ld.Get() == Syslog, "Expected Syslog")
}
func TestLogDestinationSetSad(t *testing.T){
	var ld LogDestination
	defer func(){recover()}()
	ld.Set(9)
	t.Fatal("expected error")
}