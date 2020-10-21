package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func Encrypt(signal *string, key *Key) *string {
	var err error
	var nonce []byte = make([]byte, NonceSize)
	var plaintext []byte = []byte(*signal)
	var ciphertext string
	var block cipher.Block


	block, err = aes.NewCipher((*key)[:])
	if err != nil {
		panic(err)
	}

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err)
	}

	var gcm cipher.AEAD
	gcm, err = cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	//Format:
	//	0               95                   n
	//  |-----nonce-----|-----ciphertext-----|
	//
	ciphertext = fmt.Sprintf("%x%x", nonce, gcm.Seal(nil, nonce, plaintext, nil))
	fmt.Printf("%x%x\n", nonce, gcm.Seal(nil, nonce, plaintext, nil))
	return &ciphertext

}
