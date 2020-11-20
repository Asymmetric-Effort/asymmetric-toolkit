package headers

import (
	"encoding/binary"
	"reflect"
)

func (o *dictionary.Dictionary) CreateHeader() {
	o.Header.Size = func() uint32 { // Size of the header (in bytes)
		var h HeaderDescriptor
		r := reflect.ValueOf(h)
		return uint32(binary.Size(r))
	}()
	o.Header.FormatVersion = o.Config.FormatVersion // Format version
	o.Header.ScoreVersion = o.Config.ScoreVersion   // Score version
	o.Header.WordCount = 0                          // We will populate this later.
}
