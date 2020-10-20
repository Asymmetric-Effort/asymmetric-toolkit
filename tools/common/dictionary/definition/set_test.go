package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/encryption"
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"testing"
)

func TestRecord_Set(t *testing.T) {
	var o Record
	testWord:= "test"
	passphrase:="myPassphrase"
	expectedEncodedWord:=*encryption.encrypt(&testWord, &passphrase)
	expectedId := CreateId(&testWord)
	errors.Assert(o.id == "", "Expected empty id")
	errors.Assert(o.word == "", "Expected empty word string")
	o.Set(testWord,&passphrase)
	errors.Assert(o.id==expectedId, fmt.Sprintf("Expected empty id:\n%x\n%x", o.id, expectedId))
	errors.Assert(o.word == expectedEncodedWord, fmt.Sprintf("Expected empty word string: %s", o.word))
}
