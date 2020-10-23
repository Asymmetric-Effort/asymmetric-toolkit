package definition_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source/dictionary/definition"
	"testing"
)

func TestDictionaryDefinitionConstants(t *testing.T){
	errors.Assert(definition.Delimiter == " ", "Expect single-space delimiter.")
}