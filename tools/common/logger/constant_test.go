package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestLoggerConstants(t *testing.T){
	errors.Assert(LogFormat != "", "expect non-empty string.")
	errors.Assert(logWriteInterval == 1, "Expect logWriteInterval 1")
	errors.Assert(logBufferSz == 1024, "Expect buffer size 1K")
	errors.Assert(defaultLoggerFacility == "Logger", "Expect facility string")
	errors.Assert(loggerSetupReadyMessage != "", "expect non-empty string.")
}