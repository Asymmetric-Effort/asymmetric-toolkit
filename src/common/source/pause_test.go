package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestPause(t *testing.T){
	var s Source
	errors.Assert(!s.isPaused, "Expected false isPaused as initial state")
	s.Pause()
	errors.Assert(s.isPaused,"Expect true isPaused after calling Pause()")
}