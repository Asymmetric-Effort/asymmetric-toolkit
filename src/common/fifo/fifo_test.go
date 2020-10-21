package fifo_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestFifo(t *testing.T){
	var q Fifo
	errors.Assert(q.queue == nil, "expect nil")
}