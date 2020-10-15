package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestHasData(t *testing.T){
	var s Source
	s.feed=make(chan string,10)
	errors.Assert(!s.HasData(), "Expect default to be false.")
	s.feed<-"test"
	errors.Assert(s.HasData(), "expect feed to have data.")
	<-s.feed
	errors.Assert(!s.HasData(), "Expect feed has no data.")
}