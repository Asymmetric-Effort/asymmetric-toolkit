package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strings"
	"testing"
)

func TestLevel_SetString(t *testing.T) {
	var o Level = Debug
	tests := []string{"Debug", "Info", "Warning", "Error", "Critical"}
	for _, test := range tests {
		o.SetString(test)
		expected := strings.ToLower(strings.TrimSpace(test))
		errors.Assert(o.String() == expected,
			fmt.Sprintf("Expected level mismatch (%s) != (%s)", o.String(), expected))
	}
}
