package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestEventId(t *testing.T) {
	// Core Logger operation EventId 0-999
	errors.Assert(EventStd == 0, "Expected 0")
	errors.Assert(EventInit == 1, "Expected 1")
	errors.Assert(EventTagCreate == 2, "Expected 2")
	errors.Assert(EventTagClose == 3, "Expected 3")

	// Reserved EventIds for tool and common use 1000+
}
