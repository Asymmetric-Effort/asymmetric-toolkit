package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestArgumentMap(t *testing.T){
	var o ArgumentMap
	errors.Assert(o == nil, "expect nil map")
}