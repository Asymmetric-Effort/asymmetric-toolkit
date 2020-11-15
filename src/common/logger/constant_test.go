package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLoggerConstants(t *testing.T) {
	errors.Assert(logWriteInterval == 1, "Expect logWriteInterval 1")
	errors.Assert(logBufferSz == 1024, "Expect buffer size 1K")

	errors.Assert(tagPattern == `[a-zA-Z][a-zA-Z0-9]{0,63}`, "tagPattern mismatch")
	errors.Assert(maxTagTrackerDictionarySize == 256, "maxTagTrackerDictionarySize mismatch")
}
