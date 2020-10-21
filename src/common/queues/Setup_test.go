package queues_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"testing"
)

func TestFifoSetup(t *testing.T){
	var q queues.Fifo
	errors.Assert(q.Size()==0, "Expected queue size 0")
	q.Setup(10)
	errors.Assert(q.Size()==10, "Expected queue size 10")
}