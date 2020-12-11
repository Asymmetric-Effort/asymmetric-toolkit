package dictionary

/*
	dictionary.Header::Write() writes the dictionary file header.
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func (o *Header) Write(f *os.File) (sz int64, err error) {
	var n int
	var buf bytes.Buffer
	_, err = f.Seek(0, SeekBeginning)
	if err != nil {
		return 0, err
	}
	err = binary.Write(&buf, binary.BigEndian, *o)
	errors.Assert(err == nil, fmt.Sprintf("%v", err))
	n, err = f.Write(buf.Bytes())
	return int64(n), err
}
