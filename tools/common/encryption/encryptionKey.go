package encryption

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)
const EncryptionKeyLength = 32

type EncryptionKey [32]byte

func (o *EncryptionKey) Set(passphrase *string){
	*o =sha256.Sum256([]byte(*passphrase))
}

func (o *EncryptionKey) String() string{
	return fmt.Sprintf("%02x", *o)
}

func (o *EncryptionKey) Length() int{
	return len(*o)
}

func (o *EncryptionKey) Equal(hash *[32]byte) bool {
	return bytes.Equal((*o)[:],(*hash)[:])
}