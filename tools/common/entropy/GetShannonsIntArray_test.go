package entropy

import (
	"fmt"
	"testing"
)

func TestGetShannonsIntArray(t *testing.T) {
	var tests = []struct {
		input int
		score int
	}{
		{
			0,
			0,
		},
		{
			1111111111,
			0,
		},
		{
			1223334444,
			185,
		},
		{
			9047653407,
			273,
		},
		{
			8565850718,
			238,
		},
		{
			93878445616,
			292,
		},
	}
	for i, test := range tests {
		fmt.Printf("Test %d failed with score %d (expected %d)\n", i, test.score, GetShannonsIntArray(test.input))
		if GetShannonsIntArray(test.input) > test.score {
			t.Fatalf("Test %d failed with score %d (expected %d)", i, test.score, GetShannonsIntArray(test.input))
		}
	}
}
