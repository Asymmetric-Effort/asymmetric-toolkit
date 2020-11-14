package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)
func TestLogLevel_String(t *testing.T){
	var o Level
	errors.Assert(o.String() == "critical", "expected critical")

	expected := []struct {
		d   Level
		str string
	}{
		{
			Critical, "critical",
		}, {
			Error, "error",
		},{
			Warning, "warning",
		},{
			Info, "info",
		},{
			Debug, "debug",
		},
	}
	for _,test := range expected {
		o = test.d
		errors.Assert(o.String() == test.str, "Expected "+test.str)
	}
}
