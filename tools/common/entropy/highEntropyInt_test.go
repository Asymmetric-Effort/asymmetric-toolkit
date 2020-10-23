package entropy

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"testing"
)

func TestHighEntropyInt(t *testing.T) {
	var tests = []struct {
		input int
		result bool
	}{
		{
			1111111111111111111,
			false,
		}, {
			0,
			false,
		}, {
			668108162888,
			true,
		}, {
			1234567890,
			true,
		}, {
			12345791234567890,
			true,
		}, {
			1234567890123456890,
			true,
		}, {
			1234567890123456123,
			true,
		},
	}
	for i, test := range tests {
		fmt.Printf("Test (%d) Result:%d (%t)\n", i, GetShannonsInt(test.input),HighEntropyInt(test.input))
		if test.result {
			errors.Assert(HighEntropyInt(test.input),
				fmt.Sprintf("Expected %t (%d) (%d)", HighEntropyInt(test.input), i, GetShannonsInt(test.input)))
		} else {
			errors.Assert(!HighEntropyInt(test.input),
				fmt.Sprintf("Expected %t (%d) (%d)", HighEntropyInt(test.input), i, GetShannonsInt(test.input)))
		}


	}
}
