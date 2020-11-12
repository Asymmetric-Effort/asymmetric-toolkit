package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestConfiguration(t *testing.T) {
	var o Configuration
	errors.Assert(o.Level == Critical, "Default Level is critical.")
	errors.Assert(o.Facility == "", "expect empty string.")
	errors.Assert(o.Destination == Stdout, "expect empty string.")
}
