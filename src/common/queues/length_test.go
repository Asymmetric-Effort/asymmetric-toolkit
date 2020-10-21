package queues

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLength(t *testing.T){
	var q Fifo
	errors.Assert(q.Queue == nil, "Queue expected to be nil initially.")
	q.Queue = make(chan string, 5)
	errors.Assert(q.Queue != nil, "Queue is not nil.")
	errors.Assert(q.Length()==0, "Queue length is zero.")
	q.Queue <-"1"
	errors.Assert(q.Length()==1, "Queue length is one")
	q.Queue <-"2"
	errors.Assert(q.Length()==2, "Queue length is two")
	<-q.Queue
	errors.Assert(q.Length()==1, "Queue length is one")
}