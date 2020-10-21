package definition_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary/definition"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDictionaryDefinitionConstants(t *testing.T){
	errors.Assert(definition.Delimiter == " ", "Expect single-space delimiter.")
}