package logtarget

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogDestinationSetHappy(t *testing.T){
	var ld Destination
	ld.Set(Stdout)
	errors.Assert(ld.Get() == Stdout, "Expected Stdout")
	ld.Set(File)
	errors.Assert(ld.Get() == File, "Expected File")
	ld.Set(Syslog)
	errors.Assert(ld.Get() == Syslog, "Expected Syslog")
}
func TestLogDestinationSetSad(t *testing.T){
	var ld Destination
	defer func(){recover()}()
	ld.Set(9)
	t.Fatal("expected error")
}