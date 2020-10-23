package fifo

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestFifo(t *testing.T){
	var q Fifo
	errors.Assert(q.queue == nil, "expect nil")
}