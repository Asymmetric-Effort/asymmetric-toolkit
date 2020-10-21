package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary/definition"
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary/reader"
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary/writer"
	"os"
)

type Dictionary struct {
	Read    func() string
	Write   func(s string)
	runtime struct {
		passphrase *string  //Encryption passphrase (not written to the actual file).
		fileHandle *os.File //File handle for reading/writing the actual file.
		io         struct {
			reader reader.Reader
			writer writer.Writer
		}
	}
	content struct {
		header struct {
			version           [3]byte // file format version
			reserved          [4]byte //Padding to align memory
			compressed        bool    // flag indicates whether compression is used
			created           int64   // unix timestamp
			descriptionLength uint16
			description       []byte
		}
		body struct {
			defCount    uint32 // definition count (number records in body)
			definitions []definition.Record
		}
		footer [32]byte //hash of file up to the footer.
	}
}
