package source_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"testing"
)

func TestHasData(t *testing.T){
	var s source.Source
	s.feed.Setup(10)
	errors.Assert(!s.HasData(), "Expect default to be false.")
	s.feed.Push("test")
	errors.Assert(s.HasData(), "expect feed to have data.")
	_ = s.feed.Pop()
	errors.Assert(!s.HasData(), "Expect feed has no data.")
}