package dictionaryDefinition_test

import (
	dictionaryDefinition "asymmetric-effort/asymmetric-toolkit/src/common/dictionary/definition"
	"asymmetric-effort/asymmetric-toolkit/src/common/encryption"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)
func TestRecord_Get(t *testing.T) {
	var o dictionaryDefinition.Record
	testWord:= "test"
	passphrase:="myPassphrase"
	expectedEncodedWord:= encryption.encrypt(&testWord, &passphrase)
	expectedId:= dictionaryDefinition.CreateId(&testWord)
	errors.Assert(o.id == "", "Expected empty id")
	errors.Assert(o.word == "", "Expected empty word string")
	o.id=expectedId
	o.word=*expectedEncodedWord
	errors.Assert(o.Get(&passphrase)==testWord,"Test Record::Get(): Expected test word mismatch.")
}
