package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestHasData(t *testing.T){
	var s Source
	s.feed.Setup(10)
	errors.Assert(!s.HasData(), "Expect default to be false.")
	s.feed.Push("test")
	errors.Assert(s.HasData(), "expect feed to have data.")
	_ = s.feed.Pop()
	errors.Assert(!s.HasData(), "Expect feed has no data.")
}