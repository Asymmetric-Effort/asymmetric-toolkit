package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogSettings_Get(t *testing.T) {
	var o = make(LogSettings, 1)
	o["key"] = "value"
	errors.Assert(o.Get("key") == "value", "Expected value not retrieved by Get()")
}
