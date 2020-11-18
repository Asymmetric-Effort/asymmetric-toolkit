package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDestination_String(t *testing.T) {
	expected := []struct {
		d    Destination
		str  string
		pass bool
	}{
		{
			Stdout, "stdout", true,
		}, {
			File, "file", true,
		}, {
			Syslog, "syslog", true,
		}, {
			Syslog, "Syslog", false,
		}, {
			Syslog, "Stdout", false,
		}, {
			Syslog, "STDOUT", false,
		}, {
			Syslog, "", false,
		}, {
			Syslog, "file", false,
		}, {
			Syslog, "File", false,
		}, {
			Syslog, "FILE", false,
		}, {
			Syslog, "Syslog", false,
		}, {
			Syslog, "Syslog", false,
		}, {
			Syslog, "SYSLOG", false,
		}, {
			Syslog, "bad input", false,
		},
	}
	panicOnError := func() {}
	recoverOnError := func() { recover() }
	onError := panicOnError

	testFunc := func(d Destination, s string, p bool) {
		if p {
			onError = panicOnError
		} else {
			onError = recoverOnError
		}
		defer onError()

		var o Destination = d
		errors.Assert(o.String() == s, "Expected "+s)
	}

	for _, test := range expected {
		testFunc(test.d, test.str, test.pass)
	}
}
