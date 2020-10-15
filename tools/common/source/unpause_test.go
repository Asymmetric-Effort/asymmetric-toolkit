package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestUnPause(t *testing.T) {
	var s Source
	errors.Assert(!s.isPaused, "Expected true isPaused as initial state")
	s.isPaused=true
	errors.Assert(s.isPaused, "Expected true isPaused to start test")
	s.UnPause()
	errors.Assert(!s.isPaused, "Expect false isPaused after calling UnPause()")
}
