package entropy_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/entropy"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
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
		fmt.Printf("Test (%d) Result:%d (%t)\n", i,
			entropy.GetShannonsInt(test.input), entropy.HighEntropyInt(test.input))
		if test.result {
			errors.Assert(entropy.HighEntropyInt(test.input),
				fmt.Sprintf("Expected %t (%d) (%d)",
					entropy.HighEntropyInt(test.input), i, entropy.GetShannonsInt(test.input)))
		} else {
			errors.Assert(!entropy.HighEntropyInt(test.input),
				fmt.Sprintf("Expected %t (%d) (%d)",
					entropy.HighEntropyInt(test.input), i, entropy.GetShannonsInt(test.input)))
		}


	}
}
