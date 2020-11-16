package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestMode_Get(t *testing.T) {
	tests := []struct {
		m    Mode
		pass bool
	}{
		{
			Sequence,
			true,
		}, {
			Random,
			true,
		}, {
			Dictionary,
			true,
		}, {
			3,
			false,
		}, {
			4,
			false,
		},
	}
	panicOnError := func() { /*do nothing to recover*/ }
	recoverOnError := func() { recover() }
	onError := panicOnError

	testFunc := func(pass bool, data Mode) {
		var o Mode
		o = data
		defer onError()
		errors.Assert(o.Get() == data, fmt.Sprintf("mismatched output on %d", o))
		if !pass {
			panic("Expected Error")
		}
	}
	for _, test := range tests {
		if test.pass {
			onError = panicOnError
		} else {
			onError = recoverOnError
		}
		testFunc(test.pass, test.m)
	}
}
