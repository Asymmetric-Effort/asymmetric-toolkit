package source

import "testing"

func TestMode_IsSequence(t *testing.T) {
	tests := []struct {
		value Mode
		pass  bool
	}{
		{
			value: Dictionary,
			pass:  false,
		}, {
			value: Random,
			pass:  false,
		}, {
			value: Sequence,
			pass:  true,
		},
	}
	for _, test := range tests {
		var m Mode = test.value
		if m.IsSequence() && test.pass {
			return
		} else {
			panic("IsDictionary() should be true.")
		}
	}
}