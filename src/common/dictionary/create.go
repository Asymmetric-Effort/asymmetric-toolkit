package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func (o *Dictionary) Create(filename string, description string, passphrase *string, compressed bool) {
	errors.Assert(filename != "", "Expected non-empty string as filename for dictionary")

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
		errors.Assert(length < MaxLengthDictionaryDescription, "Dictionary description must be less than 65535 characters.")
		return
	}()
	o.content.body.defCount = 0 //Empty content.
	func() {
		//
		// Setup the reader/writer
		//
		var err error
		o.runtime.fileHandle, err = os.Create(filename)
		errors.Assert(err == nil, fmt.Sprintf("Dictionary::Create() encountered error opening filename (%s): %v", filename, err))
		defer func(){
			o.runtime.io.reader.Close()
			o.runtime.io.writer.Close()
			err:=o.runtime.fileHandle.Close()
			if err != nil {
				panic(err)
			}
		}()

		o.Read = o.runtime.io.reader.Setup(o.runtime.fileHandle, []byte(*passphrase))
		o.Write = o.runtime.io.writer.Setup(o.runtime.fileHandle, []byte(*passphrase))
	}()
}
