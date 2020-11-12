package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogSettings_SetBool(t *testing.T) {
	var o LogSettings = make(LogSettings, 1)
	o.SetBool("key", true)
	errors.Assert(o["key"] == "true", "Expected bool not set")
	o.SetBool("key", false)
	errors.Assert(o["key"] == "false", "Expected bool not set")
}
