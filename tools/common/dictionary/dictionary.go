package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/dictionary/definition"
)

type Dictionary struct {
	runtime struct {
		name       *string //File name of the dictionary (not written to the actual file).
		passphrase *string //Encryption passphrase (not written to the actual file).
	}
	content struct {
		header struct {
			version           [3]byte // file format version
			compressed        bool  // flag indicates whether compression is used
			created           int64 // unix timestamp
			descriptionLength uint16
			description       []byte
		}
		body struct {
			defCount    uint32 // definition count (number records in body)
			definitions []definition.Definition
		}
		footer [32]byte //hash of file up to the footer.
	}
}
