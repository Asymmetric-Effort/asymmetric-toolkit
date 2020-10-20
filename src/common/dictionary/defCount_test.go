package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDictionary_DefCount(t *testing.T) {
	var o Dictionary
	errors.Assert(o.DefCount()==0, "Expected 0")
	o.content.body.defCount=9
	errors.Assert(o.DefCount()==9, "Expected 9")
}