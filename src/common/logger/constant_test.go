package logger_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"testing"
)

func TestLoggerConstants(t *testing.T){
	errors.Assert(logger.LogFormat != "", "expect non-empty string.")
	errors.Assert(logger.LogWriteInterval == 1, "Expect logWriteInterval 1")
	errors.Assert(logger.LogBufferSz == 1024, "Expect buffer size 1K")
	errors.Assert(logger.DefaultLoggerFacility == "Logger", "Expect facility string")
	errors.Assert(logger.LoggerSetupReadyMessage != "", "expect non-empty string.")
}