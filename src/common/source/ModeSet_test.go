package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestMode_Set(t *testing.T) {
	tests := []struct {
		m    Mode
		s string
		pass bool
	}{
		{
			Sequence,
			strModeSequence,
			true,
		}, {
			Random,
			strModeRandom,
			true,
		}, {
			Dictionary,
			strModeDictionary,
			true,
		}, {
			Sequence,
			"foo",
			false,
		}, {
			Dictionary,
				"",
			false,
		},
	}
	panicOnError := func() { /*do nothing to recover*/ }
	recoverOnError := func() { recover() }
	onError := panicOnError

	testFunc := func(pass bool, mode Mode, text string) {
		defer onError()
		var o Mode
		o.Set(mode)
		errors.Assert(o.Get() == mode, fmt.Sprintf("mismatched output on %d", o))
		errors.Assert(o.String()==text, "Mismatch text.")
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
		testFunc(test.pass, test.m, test.s)
	}
}
