package source_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"testing"
	"time"
)

func TestWaitIfPaused(t *testing.T) {
	var s source.Source
	s.IsPaused = true
	go func() {
		<-time.After(time.Second * (source.PauseCheckDelay + 5))
		t.Log("un-pausing source")
		s.IsPaused = false
	}()
	for s.IsPaused {
		t.Log("Waiting if paused.")
		errors.Assert(s.IsPaused, "Expect true")
		s.WaitIfPaused()
		errors.Assert(!s.IsPaused, "Expect false")
		t.Log("No longer paused.")
	}
}
