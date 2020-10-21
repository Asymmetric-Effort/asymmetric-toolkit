package queues_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"fmt"
	"testing"
)

func TestFifo_Pop(t *testing.T) {
	var q queues.Fifo
	const sz = 5
	fmt.Println("Initializing test")
	errors.Assert(q.IsNil(), "Queue expected to be nil initially.")
	q.Setup(sz)
	errors.Assert(!q.IsNil(), "Queue expected to be nil initially.")
	errors.Assert(q.Length()==0, "Queue length is zero.")
	fmt.Println("Pushing data into channel")
	q.Push("1")
	q.Push("2")
	fmt.Println("Starting Pop test.")
	errors.Assert(q.Length()==2, "Queue length is two")
	errors.Assert(q.Pop() == "1", "Queue Pop returned the wrong value")
	errors.Assert(q.Pop() == "2", "Queue Pop returned the wrong value")
	errors.Assert(q.Length()==0, "Queue length is 0")
}