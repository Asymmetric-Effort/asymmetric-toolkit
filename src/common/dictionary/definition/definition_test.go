package definition_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary/definition"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestRecord(t *testing.T){
	var o definition.Record
	errors.Assert(o.ID()=="", "expect empty string")
}