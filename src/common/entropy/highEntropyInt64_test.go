package entropy_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/entropy"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestHighEntropyInt64(t *testing.T) {
	var tests = []struct {
		input int64
		result bool
	}{
		{
			111111111111111111,
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
			1234567901234567890,
			true,
		}, {
			1234567901234567897,
			true,
		}, {
			1234567640134567890,
			true,
		},
	}
	for i, test := range tests {
		fmt.Printf("Test (%d) Result:%d (%t)\n", i,
			entropy.GetShannonsInt64(test.input),
			entropy.HighEntropyInt64(test.input))
		if test.result {
			errors.Assert(entropy.HighEntropyInt64(test.input),
				fmt.Sprintf("Expected %t (%d) (%d)",
					entropy.HighEntropyInt64(test.input), i,
					entropy.GetShannonsInt64(test.input)))
		} else {
			errors.Assert(!entropy.HighEntropyInt64(test.input),
				fmt.Sprintf("Expected %t (%d) (%d)",
					entropy.HighEntropyInt64(test.input), i,
					entropy.GetShannonsInt64(test.input)))
		}


	}
}
