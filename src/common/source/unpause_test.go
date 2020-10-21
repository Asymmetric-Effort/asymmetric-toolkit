package source_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"testing"
)

func TestUnPause(t *testing.T) {
	var s source.Source
	errors.Assert(!s.isPaused, "Expected true isPaused as initial state")
	s.isPaused=true
	errors.Assert(s.isPaused, "Expected true isPaused to start test")
	s.UnPause()
	errors.Assert(!s.isPaused, "Expect false isPaused after calling UnPause()")
}
