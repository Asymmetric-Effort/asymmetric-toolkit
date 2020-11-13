package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLoggerConstants(t *testing.T){



	errors.Assert(logWriteInterval == 1, "Expect logWriteInterval 1")
	errors.Assert(logBufferSz == 1024, "Expect buffer size 1K")
}
