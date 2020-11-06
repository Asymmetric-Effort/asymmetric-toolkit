package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"testing"
)

func TestLicenseHash(t *testing.T) {

	const licenseHash = "ae68d4d701e803f55cff6e18c27ba3b7f69cba59aa037745ef7ce546f5dba3f4"

	var licenseFile = "LICENSE.txt"

	hash := sha256.New()

	f, err := os.Open(licenseFile)
	if err != nil {
		t.Errorf("error opening %s: %v", licenseFile, err)
	}
	defer func() {
		e := f.Close()
		if e != nil {
			panic(e)
		}
	}()

	if _, err := io.Copy(hash, f); err != nil {
		panic(err)
	}

	actual := hex.EncodeToString(hash.Sum(nil))

	if actual != licenseHash {
		t.Errorf("Hash mismatch: %v", actual)
	}
}
