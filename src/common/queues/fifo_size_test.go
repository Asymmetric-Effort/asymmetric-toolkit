package queues_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"testing"
)

func TestFifo_Size(t *testing.T) {
	var q queues.Fifo
	errors.Assert(q.Size() == 0, "expected zero size")
	q.Setup(4)
	errors.Assert(q.Size() == 4, "expected size 4")
}
