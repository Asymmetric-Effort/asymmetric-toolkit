package fifo

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestFifo_Pop(t *testing.T) {
	var q Fifo
	errors.Assert(q.queue == nil, "Queue expected to be nil initially.")
	q.queue = make(chan string, 5)
	errors.Assert(q.queue != nil, "Queue is not nil.")
	errors.Assert(q.Length()==0, "Queue length is zero.")
	q.queue<-"1"
	q.queue<-"2"
	errors.Assert(q.Length()==2, "Queue length is two")
	errors.Assert(q.Pop() == "1", "Queue Pop returned the wrong value")
	errors.Assert(q.Pop() == "2", "Queue Pop returned the wrong value")
	errors.Assert(q.Length()==0, "Queue length is 0")
}