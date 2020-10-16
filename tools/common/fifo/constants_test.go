package fifo

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
	"time"
)

func TestFifoConstants(t *testing.T){
	errors.Assert(QueueWriteBlockDelay == time.Second *2, "Expected 2 second delay")
}