package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"strconv"
	"strings"
	"time"
)

func (o *Dictionary) Create(filename string, description string, passphrase *string, compressed bool) {
	errors.Assert(filename != "", "Expected non-empty string as filename for dictionary")
	errors.Assert(len(*passphrase) > 2047, "Expected passphrase length > 2047")
	o.runtime.name = &filename
	o.runtime.passphrase = passphrase

	o.content.header.version = func() (v [3]byte) {
		s := strings.Split(Version, ".")
		errors.Assert(len(s) == 3, "Expected 0.0.0 (three numbers) in version string.")
		for i := 0; i <= 3; i++ {
			n, err := strconv.Atoi(s[i])
			errors.Assert(err == nil, "Unexpected error in Dictionary::Create() parsing version string")
			v[i] = byte(n)
		}
		return
	}()
	o.content.header.compressed = compressed
	o.content.header.created = time.Now().Unix()
	o.content.header.description = []byte(description)
	o.content.header.descriptionLength = func() (length uint16) {
		length = uint16(len(o.content.header.description))
		errors.Assert(length < 65535, "Dictionary description must be less than 65535 characters.")
		return
	}()
	o.content.body.defCount = 0 //Empty content.
}