package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/encryption"
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
)

func (o *Record) Set(s string, passphrase *string) {
	errors.Assert(o.id == "", "Record::Set(): can only be used once to set the word value.")
	o.id = CreateId(&s)
	o.word = *encryption.encrypt(&s, passphrase)
}
