package entropy_test

import (
	"fmt"
	"testing"
)

func TestGetShannonsInt64Array(t *testing.T) {
	var tests = []struct {
		input int64
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
			8565850718799581350,
			288,
		},
		{
			1938778995290015530,
			291,
		},
	}
	for i, test := range tests {
		fmt.Printf("Test %d failed with score %d (expected %d)\n", i, test.score, GetShannonsInt64(test.input))
		if GetShannonsInt64(test.input) > int64(test.score) {
			t.Fatalf("Test %d failed with score %d (expected %d)", i, test.score, GetShannonsInt64(test.input))
		}
	}
}
