package writer

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func (o *Writer) Setup(file *os.File) (write func(s string)) {
	errors.Assert(file != nil, "Nil file handle encountered in dictionary Writer::Setup()")
	writer := bufio.NewWriter(o.file)
	/*
		Create the write() function.
	*/
	write = func(s string) {
		encodedString := base64.StdEncoding.EncodeToString([]byte(s))
		out := fmt.Sprintf("%d %s 0000000000 	0 0", len(encodedString), encodedString)
		i, err := writer.WriteString(out)
		errors.Assert(err == nil, fmt.Sprintf("Encountered error writing to dictionary.  error:%v", err))
		errors.Assert(i == len(out), fmt.Sprintf("Expected to write (%d) to dictionary.  Wrote (%d)", len(out), i))
	}
	return
}
