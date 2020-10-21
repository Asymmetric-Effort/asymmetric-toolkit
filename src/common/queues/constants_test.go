package queues_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"testing"
	"time"
)

func TestFifoConstants(t *testing.T){
	errors.Assert(queues.QueueWriteBlockDelay == time.Second *2, "Expected 2 second delay")
}