package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestArgumentType(t *testing.T) {
	func() {
		tests := []struct {
			Type  ArgumentType
			Value int
		}{
			{
				None, 0,
			},
			{
				String, 1,
			},
			{
				Integer, 2,
			},
			{
				Float, 3,
			},
			{
				Boolean, 4,
			},
		}
		for _, test := range tests {
			errors.Assert(int(test.Type) == test.Value, "Expected value mismatch")
		}
	}()
}
