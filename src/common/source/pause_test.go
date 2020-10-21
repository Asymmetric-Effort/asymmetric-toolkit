package source_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"testing"
)

func TestPause(t *testing.T){
	var s source.Source
	errors.Assert(!s.IsPaused, "Expected false isPaused as initial state")
	s.Pause()
	errors.Assert(s.IsPaused,"Expect true isPaused after calling Pause()")
}