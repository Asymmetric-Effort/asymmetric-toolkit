package queues_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"testing"
)

func TestFifo(t *testing.T){
	var q queues.Fifo
	errors.Assert(q.Queue == nil, "expect nil")
}