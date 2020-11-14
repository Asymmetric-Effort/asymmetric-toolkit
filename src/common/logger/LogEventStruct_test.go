package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
	"time"
)

func TestLogEventStruct(t *testing.T) {
	var o LogEventStruct
	errors.Assert(o.eventId == 0, "expect 0")
	errors.Assert(o.time == time.Time{}, "expect 0")
	errors.Assert(o.level == Critical, "expect critical (0)")
	errors.Assert(o.tags == nil, "expect nil")
	errors.Assert(o.message == "", "expect nil")
}
