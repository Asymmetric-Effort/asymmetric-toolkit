package encryption

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"crypto/aes"
	"crypto/cipher"

	//"crypto/aes"
	//"crypto/cipher"
	//"encoding/hex"
	"encoding/hex"
	"fmt"
)

func Decrypt(signal *string, key *Key) *string {
	var err error
	var nonce *[]byte
	var plainText []byte
	var cipherText *[]byte
	var block cipher.Block

	cipherText, nonce = func() (*[]byte, *[]byte) {
		//Parse the original signal into the nonce and ciphertext.
		//Format:
		//	0               95                   n
		//  |-----nonce-----|-----ciphertext-----|
		//
		sbytes, err := hex.DecodeString(*signal)
		errors.Assert(err == nil,
			fmt.Sprintf("Failed to decode ciphertext string. "+
				"Error:%v\n"+
				"signal:%s",
				*signal, err))
		n := sbytes[0:nonceSize]
		c := sbytes[nonceSize:]

		return &c, &n
	}()

	fmt.Println(cipherText, nonce)

	block, err = aes.NewCipher((*key)[:])
	errors.Assert(err == nil, fmt.Sprintf("Failed to generate the cipher block. Error:%v", err))

	gcm, err := cipher.NewGCM(block)
	errors.Assert(err == nil, fmt.Sprintf("Failed to create new GCM. Error:%v", err))

	plainText, err = gcm.Open(nil, *nonce, *cipherText, nil)
	errors.Assert(err == nil, fmt.Sprintf("Decryption failed. Error:%v", err))
	s := string(plainText)

	return &s
}
