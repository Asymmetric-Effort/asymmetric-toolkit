package queues_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"testing"
)

func TestLength(t *testing.T){
	var q queues.Fifo
	q.Setup(5)
	errors.Assert(q.Length()==0, "queue length is zero.")
	q.Push("1")
	errors.Assert(q.Length()==1, "queue length is one")
	q.Push("2")
	errors.Assert(q.Length()==2, "queue length is two")
	q.Pop()
	errors.Assert(q.Length()==1, "queue length is one")
}