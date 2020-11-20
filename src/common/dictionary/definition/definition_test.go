package definition

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDefinition(t *testing.T) {
	var o Descriptor
	errors.Assert(o.Id == 0, "expect 0")
	errors.Assert(o.Word == "", "expect empty string")
	errors.Assert(o.Score == 0, "expect 0")
	errors.Assert(o.Created == 0, "expect 0")
	errors.Assert(o.LastHit == 0, "expect 0")
	errors.Assert(&o.Tags != nil, "expect non-nil pointer")
	errors.Assert(o.Hits == 0, "expect 0")
	errors.Assert(o.Miss == 0, "expect 0")
}
