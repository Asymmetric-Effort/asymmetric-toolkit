package queues_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"testing"
)

func TestFifoSetup(t *testing.T){
	var q queues.Fifo
	q.Setup(10)
}