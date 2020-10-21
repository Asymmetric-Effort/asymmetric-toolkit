package queues_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"testing"
)

func TestFifo_IsNil(t *testing.T) {
	var q queues.Fifo
	errors.Assert(q.IsNil(), "Expected true")
	q.Setup(1)
	errors.Assert(!q.IsNil(), "Expected false")
}