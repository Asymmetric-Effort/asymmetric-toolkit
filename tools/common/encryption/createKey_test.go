package encryption

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"testing"
)

func TestCreateKey(t *testing.T) {
	passphrases := []string{
		"6368616e676520746869732070617373",
		"password01",
		"cpe-1704-tks",
	}
	for i, passphrase := range passphrases {
		fmt.Println("test:", i)
		key := createKey(&passphrase)
		errors.Assert(len(*key) == 32, fmt.Sprintf("Test #%d: Expected 32-byte key.", i))
	}
}

