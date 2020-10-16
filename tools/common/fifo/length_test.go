package fifo

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestLength(t *testing.T){
	var q Fifo
	errors.Assert(q.queue == nil, "Queue expected to be nil initially.")
	q.queue = make(chan string, 5)
	errors.Assert(q.queue != nil, "Queue is not nil.")
	errors.Assert(q.Length()==0, "Queue length is zero.")
	q.queue<-"1"
	errors.Assert(q.Length()==1, "Queue length is one")
	q.queue<-"2"
	errors.Assert(q.Length()==2, "Queue length is two")
	<-q.queue
	errors.Assert(q.Length()==1, "Queue length is one")
}