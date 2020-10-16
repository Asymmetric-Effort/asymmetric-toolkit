package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestSourceHappy(t *testing.T) {
	var s Source
	errors.Assert(s.config == nil, "Expect nil")
	errors.Assert(s.allowedChars == nil, "expect nil")
	errors.Assert(s.sourceBufferSz == 0, "expected 0")
	errors.Assert(!s.isPaused, "expect false")
	errors.Assert(s.feed.Length() == 0, "expect Fifo Queue Length 0")
	errors.Assert(s.counter == 0, "expected 0 counter")
	errors.Assert(s.logger == nil, "expected nil logger")
}

