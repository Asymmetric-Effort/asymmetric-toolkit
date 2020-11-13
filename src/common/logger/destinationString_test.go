package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDestination_String(t *testing.T) {
	var o Destination

	errors.Assert(o.String() == "stdout", "expected stdout")

	expected := []struct {
		d   Destination
		str string
	}{
		{
			Stdout, "stdout",
		}, {
			File, "file",
		}, {
			Syslog, "syslog",
		},
	}
	for _,test := range expected {
		o = test.d
		errors.Assert(o.String() == test.str, "Expected "+test.str)

	}
}
