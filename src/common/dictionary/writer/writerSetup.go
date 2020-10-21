package DictionaryWriter

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"io"
	"os"
)

func (o *Writer) Setup(file *os.File, passphrase []byte) (write func(s string)) {
	errors.Assert(file != nil, "Nil file handle encountered in dictionary Writer::Setup()")
	o.file = file
	writer := bufio.NewWriter(file)
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
	return func(s string) {
		var err error
		nonce := make([]byte, gcm.NonceSize())
		_, err = io.ReadFull(rand.Reader, nonce)
		errors.Assert(err == nil, fmt.Sprintf("%v", err))
		_, err = writer.Write(
			gcm.Seal(
				nonce,
				nonce,
				[]byte(base64.StdEncoding.EncodeToString([]byte(s))),
				nil))
		assert.Error(err==nil, "Errorf calling write() operation in DictionaryWriter.Write()")
		err = writer.Flush()
		errors.Assert(err == nil, fmt.Sprintf("DictionaryWriter::Write() failed to flush.  Errorf: %v", err))
		errors.Assert(err == nil, fmt.Sprintf("%v", err))
	}
}
