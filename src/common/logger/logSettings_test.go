package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogSettings(t *testing.T) {
	var o LogSettings
	o = make(LogSettings, 1)
	o["test"] = "value"
	errors.Assert(o["test"] == "value", "Expected initial value not found.")
}
