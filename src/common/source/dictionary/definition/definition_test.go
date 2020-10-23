package definition_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source/dictionary/definition"
	"testing"
)

func TestRecord(t *testing.T){
	var o definition.Record
	errors.Assert(o.ID()=="", "expect empty string")
}