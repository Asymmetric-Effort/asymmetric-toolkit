package destination

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogDestinationGetHappy(t *testing.T){
	var ld LogDestination
	ld= Stdout
	errors.Assert(ld.Get() == Stdout, "Expected Stdout")
	ld= File
	errors.Assert(ld.Get() == File, "Expected File")
	ld= Syslog
	errors.Assert(ld.Get() == Syslog, "Expected Syslog")
}
func TestLogDestinationGetSad(t *testing.T){
	var ld LogDestination = 9
	defer func(){recover()}()
	_ = ld.Get()
	t.Fatal("expected error")
}