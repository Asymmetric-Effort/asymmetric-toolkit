package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestDictionaryDefinitionConstants(t *testing.T){
	errors.Assert(delimiter == " ", "Expect single-space delimiter.")
}