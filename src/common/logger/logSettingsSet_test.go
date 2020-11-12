package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogSettings_Set(t *testing.T) {
	var o LogSettings = make(LogSettings,1)
	o.Set("key","value")
	errors.Assert(o["key"]== "value", "Expected value not set.")
}