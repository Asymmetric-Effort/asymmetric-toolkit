package fifo

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestFifo_CloseNil(t *testing.T) {
	var q Fifo
	errors.Assert(q.queue==nil, "Expect initial queue to be nil")
	defer func(){recover()}()
	q.Close()
	t.Error("Expected error on closing nil queue")
}

func TestFifo_CloseHappy(t *testing.T) {
	const sz = 10
	var q Fifo
	q.queue = make(chan string, sz)
	q.queue<-"1"
	errors.Assert(q.Pop()=="1","Expected popped value (1)")
	q.queue<-"2"
	q.Close()
	errors.Assert(q.Pop()=="2","Expected popped value (2)")
	defer func(){recover()}()
	q.queue<-"3"
	t.Error("Expected error pushing to closed queue.")
}