package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogSettings_SetInt(t *testing.T) {
	var o = make(LogSettings,1)
	o.SetInt("key",1)
	errors.Assert(o["key"]=="1", "Expected value not set.")
}