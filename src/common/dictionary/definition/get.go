package dictionaryDefinition

import "asymmetric-effort/asymmetric-toolkit/src/common/encryption"

func (o *Record) Get(key *encryption.Key) string {
	return *encryption.Decrypt(&o.word,key)
}
