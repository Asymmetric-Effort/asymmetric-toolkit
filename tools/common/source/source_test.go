package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
	"time"
)

func TestSourceHappy(t *testing.T) {
	var s Source
	errors.Assert(s.config == nil, "Expect nil")
	errors.Assert(s.allowedChars == nil, "expect nil")
	errors.Assert(s.sourceBufferSz == 0, "expected 0")
	errors.Assert(!s.isPaused, "expect false")
	errors.Assert(s.feed == nil, "expect nil channel")
	errors.Assert(s.counter == 0, "expected 0 counter")
	errors.Assert(s.logger == nil, "expected nil logger")
}

func testTimeout(t *testing.T, to time.Duration) {
	ticker := time.NewTicker(to * time.Second)
	defer func() { ticker.Stop() }()
	go func() { //Timeout
		for _ = range ticker.C {
			t.Fatal("Timeout exceeded")
		}
	}()
}
