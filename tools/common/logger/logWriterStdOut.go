package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"bufio"
	"fmt"
	"os"
)

func (o *Logger) logWriterStdOut(msg *string) {
	writer := bufio.NewWriter(os.Stdout)
	_, err := writer.WriteString(*msg)
	errors.Assert(err == nil, fmt.Sprintf("%v",err))
	defer func() {
		err := writer.Flush()
		errors.Assert(err == nil, fmt.Sprintf("%v",err))
	}()
}
