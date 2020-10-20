package entropy

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestHighEntropy(t *testing.T) {
	var tests = []struct {
		input string
		result bool
	}{
		{
			"1111111111111111111111",
			false,
		}, {
			"",
			false,
		}, {
			"0",
			false,
		}, {
			"false",
			true,
		}, {
			"668108162888",
			true,
		}, {
			"1234567890",
			true,
		}, {
			"12345678901234567890",
			true,
		}, {
			"123456789012345678901234567890",
			true,
		}, {
			"12345678901234567890123456789012345678901234567890",
			true,
		},
	}
	for i, test := range tests {
		fmt.Printf("Test (%d) Result:%d (%t)\n", i, GetShannons(test.input), HighEntropy(test.input))
		if test.result {
			errors.Assert(HighEntropy(test.input),
				fmt.Sprintf("Expected %t (%d) (%d)", HighEntropy(test.input), i, GetShannons(test.input)))
		} else {
			errors.Assert(!HighEntropy(test.input),
				fmt.Sprintf("Expected %t (%d) (%d)", HighEntropy(test.input), i, GetShannons(test.input)))
		}


	}
}

/*
	1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 => 4.5849 shannons
	1 2 3 4 5 6 1 2 3 4 5 6 1 2 3 4 5 6 1 2 3 4 5 6 => 4.3983 shannons
 */