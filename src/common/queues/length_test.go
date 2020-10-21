package queues_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"testing"
)

func TestLength(t *testing.T){
	var q queues.Fifo
	const sz = 5
	errors.Assert(q.IsNil(), "Queue expected to be nil initially.")
	q.Setup(sz)
	errors.Assert(!q.IsNil(), "Queue expected to be nil initially.")
	errors.Assert(q.Length()==0, "Queue length is zero.")
	q.Push("1")
	errors.Assert(q.Length()==1, "Queue length is one")
	q.Push("2")
	errors.Assert(q.Length()==2, "Queue length is two")
	q.Pop()
	errors.Assert(q.Length()==1, "Queue length is one")
}