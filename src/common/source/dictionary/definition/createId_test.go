package definition_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	dictionaryDefinition "asymmetric-effort/asymmetric-toolkit/src/common/source/dictionary/definition"
	"testing"
)

func TestCreateId(t *testing.T) {
	testWord:="test"
	hash:= dictionaryDefinition.CreateId(&testWord)
	errors.Assert(hash=="74657374e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","Hash mismatch")
}