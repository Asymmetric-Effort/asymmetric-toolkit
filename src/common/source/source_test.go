package source_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"testing"
)

func TestSourceHappy(t *testing.T) {
	var s source.Source
	errors.Assert(s.Config == nil, "Expect nil")
	errors.Assert(s.AllowedChars == nil, "expect nil")
	errors.Assert(s.SourceBufferSz == 0, "expected 0")
	errors.Assert(!s.IsPaused, "expect false")
	errors.Assert(s.Feed.Length() == 0, "expect Fifo Queue Length 0")
	errors.Assert(s.Counter == 0, "expected 0 counter")
	errors.Assert(s.Logger == nil, "expected nil logger")
}

