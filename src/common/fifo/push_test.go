package fifo

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestFifo_Push(t *testing.T) {
	var q Fifo
		q.sz=2
	errors.Assert(q.queue == nil, "Queue expected to be nil initially.")
	errors.Assert(q.sz==2, "Queue Size == 2")
	q.queue = make(chan string, q.sz)
	errors.Assert(q.queue != nil, "Queue is not nil.")
	count:=0
	go func() {
		errors.Assert(q.Length() == 0, "Queue length is zero.")
		for i:=1;i<=4;i++ {
			t.Logf("Pushing %d to queue\n",i)
			q.queue <- strconv.Itoa(i)
			count++
		}
		errors.Assert(count == 4, "Expect 4 processed  push() requests")
	}()
	<-time.After(5*time.Second)
	t.Logf("Starting to collect\n")
	errors.Assert(count == 2, "Expect 2 processed push() requests on the receiver.")

	t.Logf("Step 1 Length:%d\n",q.Length())
	errors.Assert(q.Length() == 2, fmt.Sprintf("Step 1. Queue length expects 2 (capacity). Encountered:%d",q.Length()))
	<-time.After(2*time.Second)

	t.Logf("Step 2 Length:%d\n",q.Length())
	errors.Assert(q.Length() == 2, fmt.Sprintf("Step 2. Queue length expects 2. Encountered:%d",q.Length()))
	errors.Assert(q.Pop() == "1", "Queue Pop returned the wrong value")

	t.Logf("Step 3 Length:%d\n",q.Length())
	errors.Assert(q.Length() == 2,fmt.Sprintf("Step 3. Queue length expects 2. Encountered:%d",q.Length()))
	errors.Assert(q.Pop() == "2", "Queue Pop returned the wrong value")

	t.Logf("Step 4 Length:%d\n",q.Length())
	errors.Assert(q.Length() == 1, fmt.Sprintf("Step 4. Queue length expects 1. Encountered:%d",q.Length()))
	errors.Assert(q.Pop() == "3", "Queue Pop returned the wrong value")

	t.Logf("Step 5 Length:%d\n",q.Length())
	errors.Assert(q.Length() == 0, fmt.Sprintf("Step 5. Queue length expects 0. Encountered:%d",q.Length()))
	errors.Assert(q.Pop() == "4", "Queue Pop returned the wrong value")
}
