package dictionary

import (
	"os"
)

func (o *Writer) Setup(file *os.File) func() *string {
	return func() *string {
		return nil
	}
}