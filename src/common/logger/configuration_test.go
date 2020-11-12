package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
	"testing"
)

func TestConfiguration(t *testing.T) {
	var o Configuration
	errors.Assert(o.level == logLevel.Critical, "Default level is critical.")
	errors.Assert(o.facility == "", "expect empty string.")
	errors.Assert(o.target == "", "expect empty string.")
}
