package dictionaryDefinition_test

import (
	dictionaryDefinition "asymmetric-effort/asymmetric-toolkit/src/common/dictionary/definition"
	"asymmetric-effort/asymmetric-toolkit/src/common/encryption"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestRecord_Set(t *testing.T) {
	var o dictionaryDefinition.Record
	testWord:= "test"
	passphrase:="myPassphrase"
	expectedEncodedWord:=*encryption.encrypt(&testWord, &passphrase)
	expectedId := dictionaryDefinition.CreateId(&testWord)
	errors.Assert(o.id == "", "Expected empty id")
	errors.Assert(o.word == "", "Expected empty word string")
	o.Set(testWord,&passphrase)
	errors.Assert(o.id==expectedId, fmt.Sprintf("Expected empty id:\n%x\n%x", o.id, expectedId))
	errors.Assert(o.word == expectedEncodedWord, fmt.Sprintf("Expected empty word string: %s", o.word))
}
