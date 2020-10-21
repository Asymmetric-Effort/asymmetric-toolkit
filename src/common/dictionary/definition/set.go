package definition

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/encryption"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
)

func (o *Record) Set(s string, key *encryption.Key) {
	errors.Assert(o.id == "", "Record::Set(): can only be used once to set the word value.")
	o.id = CreateID(&s)
	o.word = *encryption.Encrypt(&s, key)
}
