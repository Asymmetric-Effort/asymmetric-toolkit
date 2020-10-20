package encryption

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/random"
	"golang.org/x/crypto/bcrypt"
)

func createKey(passphrase *string) *[]byte {
	/*
		This is not safe or production ready.  I'm doing this to get started.
	 */
	key, _ := bcrypt.GenerateFromPassword(
		[]byte(*passphrase),
		random.Number(bcrypt.DefaultCost,bcrypt.MaxCost))

	return &key
}