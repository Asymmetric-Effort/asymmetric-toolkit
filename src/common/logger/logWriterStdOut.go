package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bufio"
	"fmt"
	"os"
)

func (o *Logger) logWriterStdOut(msg *[]byte) {
	writer := bufio.NewWriter(os.Stdout)
	_, err:=writer.Write(*msg)
	errors.Assert(err == nil, fmt.Sprintf("%v",err))
	_ = writer.Flush()
}
