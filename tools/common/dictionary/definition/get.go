package dictionaryDefinition

import "asymmetric-effort/asymmetric-toolkit/tools/common/encryption"

func (o *Record) Get(passphrase *string) string {
	return *encryption.decrypt(&o.word,passphrase)
}
