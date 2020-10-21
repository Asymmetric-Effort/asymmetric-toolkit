package fifo_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
	"time"
)

func TestFifoConstants(t *testing.T){
	errors.Assert(QueueWriteBlockDelay == time.Second *2, "Expected 2 second delay")
}