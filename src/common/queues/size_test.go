package queues_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"testing"
)

func TestFifo_Size(t *testing.T) {
	var o queues.Fifo
	errors.Assert(o.Size()==0,"Expected size ==0")
	o.Setup(2)
	errors.Assert(o.Size()==2,"Expected size ==2")
}