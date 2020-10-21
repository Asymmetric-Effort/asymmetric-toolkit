package queues_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"testing"
)

func TestFifo_CloseNil(t *testing.T) {
	var q queues.Fifo
	defer func() { recover() }()
	q.Close()
	t.Error("Expected error on closing nil queue")
}

func TestFifo_CloseHappy(t *testing.T) {
	const sz = 10
	var q queues.Fifo
	q.Setup(sz)
	q.Push("1")
	errors.Assert(q.Pop() == "1", "Expected popped value (1)")
	q.Push("2")
	q.Close()
	errors.Assert(q.Pop() == "2", "Expected popped value (2)")
	defer func() { recover() }()
	q.Push("3")
	t.Error("Expected error pushing to closed queue.")
}
