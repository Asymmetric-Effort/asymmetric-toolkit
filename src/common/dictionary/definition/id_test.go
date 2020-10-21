package definition

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRecord_Id(t *testing.T) {
	var o Record
	testWord := "test"
	expectedID:=fmt.Sprintf("%x",sha256.Sum256([]byte(testWord)))
	expectedEncodedWord := base64.StdEncoding.EncodeToString([]byte(testWord))
	errors.Assert(o.ID() == "", "Expected empty id")
	errors.Assert(o.String() == "", "Expected empty word string")
	o.id = CreateId(&testWord)
	o.word = expectedEncodedWord
	errors.Assert(o.ID() == expectedID, "Test Record::Get(): expectedId.")
}
