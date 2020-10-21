package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
	"time"
)

func TestWaitIfPaused(t *testing.T) {
	var s Source
	s.isPaused = true
	go func() {
		<-time.After(time.Second * (pauseCheckDelay + 5))
		t.Log("un-pausing source")
		s.isPaused = false
	}()
	for s.isPaused {
		t.Log("Waiting if paused.")
		errors.Assert(s.isPaused, "Expect true")
		s.WaitIfPaused()
		errors.Assert(!s.isPaused, "Expect false")
		t.Log("No longer paused.")
	}
}
