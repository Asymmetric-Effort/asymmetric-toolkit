package logger

import (
	"fmt"
)

func (o *Logger) logWriterStdOut(msg *string) {
	fmt.Print(*msg)
}
