package definition_test

import (
	dictionaryDefinition "asymmetric-effort/asymmetric-toolkit/src/common/dictionary/definition"
	"asymmetric-effort/asymmetric-toolkit/src/common/encryption"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestRecord_Set(t *testing.T) {
	var o dictionaryDefinition.Record
	var key encryption.Key
	testWord:= "test"
	passphrase:="myPassphrase"
	key.Set(&passphrase)

	expectedEncodedWord:=*encryption.Encrypt(&testWord, &key)
	expectedId := dictionaryDefinition.CreateId(&testWord)
	errors.Assert(o.ID() == "", "Expected empty id")
	errors.Assert(o.String(&key) == "", "Expected empty word string")
	o.Set(testWord,&key)
	errors.Assert(o.ID()==expectedId,
		fmt.Sprintf("Expected empty id:\n%x\n%x", o.ID(), expectedId))
	errors.Assert(o.String(&key) == expectedEncodedWord,
		fmt.Sprintf("Expected empty word string: %s", o.String(&key)))
}
