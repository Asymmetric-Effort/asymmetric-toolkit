package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/encryption"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)
func TestRecord_Get(t *testing.T) {
	var o Record
	testWord:= "test"
	passphrase:="myPassphrase"
	expectedEncodedWord:= encryption.encrypt(&testWord, &passphrase)
	expectedId:= CreateId(&testWord)
	errors.Assert(o.id == "", "Expected empty id")
	errors.Assert(o.word == "", "Expected empty word string")
	o.id=expectedId
	o.word=*expectedEncodedWord
	errors.Assert(o.Get(&passphrase)==testWord,"Test Record::Get(): Expected test word mismatch.")
}
