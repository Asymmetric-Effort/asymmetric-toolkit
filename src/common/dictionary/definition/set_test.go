package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/encryption"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestRecord_Set(t *testing.T) {
	var o Record
	var key encryption.Key
	testWord:= "test"
	passphrase:="myPassphrase"
	key.Set(&passphrase)
	expectedEncodedWord:=*encryption.Encrypt(&testWord, &key)
	expectedId := CreateId(&testWord)
	errors.Assert(o.id == "", "Expected empty id")
	errors.Assert(o.word == "", "Expected empty word string")
	o.Set(testWord,&key)
	errors.Assert(o.id==expectedId, fmt.Sprintf("Expected empty id:\n%x\n%x", o.id, expectedId))
	errors.Assert(o.word == expectedEncodedWord, fmt.Sprintf("Expected empty word string: %s", o.word))
}
