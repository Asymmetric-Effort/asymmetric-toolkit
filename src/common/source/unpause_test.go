package source_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"testing"
)

func TestUnPause(t *testing.T) {
	var s source.Source
	errors.Assert(!s.IsPaused, "Expected true isPaused as initial state")
	s.IsPaused=true
	errors.Assert(s.IsPaused, "Expected true isPaused to start test")
	s.UnPause()
	errors.Assert(!s.IsPaused, "Expect false isPaused after calling UnPause()")
}
