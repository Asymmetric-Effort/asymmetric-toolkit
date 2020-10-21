package dictionaryDefinition_test

import (
	dictionaryDefinition "asymmetric-effort/asymmetric-toolkit/src/common/dictionary/definition"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDictionaryDefinitionConstants(t *testing.T){
	errors.Assert(dictionaryDefinition.Delimiter == " ", "Expect single-space Delimiter.")
}