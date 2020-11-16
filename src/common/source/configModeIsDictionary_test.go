package source

import (
	"testing"
)

func TestMode_IsDictionary(t *testing.T) {
	tests := []struct {
		value Mode
		pass  bool
	}{
		{
			value: Dictionary,
			pass:  true,
		}, {
			value: Random,
			pass:  false,
		}, {
			value: Sequence,
			pass:  false,
		},
	}
	for _, test := range tests {
		var m Mode = test.value
		if m.IsDictionary() && test.pass {
			return
		} else {
			panic("IsDictionary() should be true.")
		}
	}
}
