package dictionary

/*
	dictionary.Header::Read()
*/
import (
	"encoding/binary"
	"os"
)

func (o *Header) Read(f *os.File) (err error) {
	_, err = f.Seek(0, SeekBeginning)
	if err != nil {
		return err
	}
	return binary.Read(f, binary.BigEndian, o)
}
