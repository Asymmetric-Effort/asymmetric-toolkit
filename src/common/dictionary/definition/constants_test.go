package definition

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDictionaryDefinitionConstants(t *testing.T){
	errors.Assert(delimiter == " ", "Expect single-space delimiter.")
}