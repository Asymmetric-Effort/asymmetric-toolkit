package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestHasData(t *testing.T){
	var s Source
	s.queue.Setup(10)
	errors.Assert(!s.HasData(), "Expect default to be false.")
	s.queue.Push("test")
	errors.Assert(s.HasData(), "expect feed to have data.")
	_ = s.queue.Pop()
	errors.Assert(!s.HasData(), "Expect feed has no data.")
}