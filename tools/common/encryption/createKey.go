package encryption

import (
	"crypto/sha256"
)

func createKey(passphrase *string) *[32]byte {
	/*
		This is not safe or production ready.  I'm doing this to get started.
	*/
	key := sha256.Sum256([]byte(*passphrase))
	return &key
}
