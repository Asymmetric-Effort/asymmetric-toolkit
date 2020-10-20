package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestRecord(t *testing.T){
	var o Record
	errors.Assert(o.id=="", "expect empty string")
	errors.Assert(o.word=="", "expect empty word string")
}