package dictionaryDefinition_test

import (
	dictionaryDefinition "asymmetric-effort/asymmetric-toolkit/src/common/dictionary/definition"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestRecord(t *testing.T){
	var o dictionaryDefinition.Record
	errors.Assert(o.id=="", "expect empty string")
	errors.Assert(o.word=="", "expect empty word string")
}