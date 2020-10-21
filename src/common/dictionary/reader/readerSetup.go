package DictionaryReader

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func (o *Reader) Setup(file *os.File, passphrase []byte) (read func() string) {
	errors.Assert(file != nil, "Reader::Setup() encountered nil file handle")
	o.file = file
	scanner := bufio.NewScanner(file)
	encryptionKey := func() []byte {
		//Generate a sha256 hash of our passphrase.
		key := []byte(passphrase)
		hash := sha256.New()
		return hash.Sum(key)
	}()
	c, err := aes.NewCipher(encryptionKey)
	errors.Assert(err == nil, fmt.Sprintf("%v", err))
	gcm, err := cipher.NewGCM(c)
	errors.Assert(err == nil, fmt.Sprintf("%v", err))
	return func() string {
		var err error
		var line []byte
		nonce := make([]byte, gcm.NonceSize())
		_, err = io.ReadFull(rand.Reader, nonce)
		errors.Assert(err == nil, fmt.Sprintf("%v", err))
		if scanner.Scan() {
			line, err = base64.StdEncoding.DecodeString(
				func() string {
					ciphertext := scanner.Text()
					nonceSize := gcm.NonceSize()
					errors.Assert(len(ciphertext) < nonceSize, fmt.Sprintf("%v", err))
					nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
					plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
					errors.Assert(err == nil, fmt.Sprintf("DictionaryReader::Read(): Bad format.  Error: %v", err))
					return string(plaintext)
				}())
		}
		return string(line)
	}
}
