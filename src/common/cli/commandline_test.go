package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestCommandLine(t *testing.T){
	var o CommandLine
	errors.Assert(o.Arguments==nil,"Expected nil Arguments map.")
	o.Arguments = make(map[ArgumentFlag]*Argument)
	errors.Assert(o.Arguments[noFlag]==nil, "nil Argument")
}