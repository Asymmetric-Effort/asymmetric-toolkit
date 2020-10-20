package encryption

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

const KeyLength = 32

type Key [32]byte

func (o *Key) Set(passphrase *string){
	*o =sha256.Sum256([]byte(*passphrase))
}

func (o *Key) String() string{
	return fmt.Sprintf("%02x", *o)
}

func (o *Key) Length() int{
	return len(*o)
}

func (o *Key) Equal(hash *[32]byte) bool {
	return bytes.Equal((*o)[:],(*hash)[:])
}
