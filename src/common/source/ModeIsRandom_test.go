package source

import "testing"

func TestMode_IsRandom(t *testing.T) {
	tests := []struct {
		value Mode
		pass  bool
	}{
		{
			value: Dictionary,
			pass:  false,
		}, {
			value: Random,
			pass:  true,
		}, {
			value: Sequence,
			pass:  false,
		},
	}
	for _, test := range tests {
		var m Mode = test.value
		if m.IsRandom() && test.pass {
			return
		} else {
			panic("IsDictionary() should be true.")
		}
	}
}
