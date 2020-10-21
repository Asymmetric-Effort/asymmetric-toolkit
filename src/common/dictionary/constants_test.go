package dictionary_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDictionaryConstants(t *testing.T){
	errors.Assert(dictionary.Version == "0.0.1", "expected version 0.0.1")
	errors.Assert(dictionary.MaxLengthDictionaryDescription == uint16(65535), "expected 65535 (uint16)")
}