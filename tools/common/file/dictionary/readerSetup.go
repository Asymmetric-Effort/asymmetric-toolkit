package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"bufio"
	"os"
)

func (o *Reader) Setup(file *os.File) func() *string {
	errors.Assert(file != nil, "Reader::Setup() encountered nil file handle")
	o.file = file
	scanner := bufio.NewScanner(file)

	return func() *string {
		if scanner.Scan() {
			line := scanner.Text()
			//ToDo: We need a dictionary reader to handle the dictionary format to allow
			//      more intelligent dictionaries.
			//
			//      Dictionary record: <orderId: int64> <id: int64> <word: base64string>
			//
			return &line
		}
		return nil
	}
}