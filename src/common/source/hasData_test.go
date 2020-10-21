package source_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"testing"
)

func TestHasData(t *testing.T){
	var s source.Source
	s.Feed.Setup(10)
	errors.Assert(!s.HasData(), "Expect default to be false.")
	s.Feed.Push("test")
	errors.Assert(s.HasData(), "expect feed to have data.")
	_ = s.Feed.Pop()
	errors.Assert(!s.HasData(), "Expect feed has no data.")
}