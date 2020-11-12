package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogSettings_GetBool(t *testing.T) {
	var o LogSettings = make(LogSettings, 1)
	o["key"] = "true"
	errors.Assert(o.GetBool("key"), "Boolean key expected true")
}
