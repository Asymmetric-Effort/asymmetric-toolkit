package errors

import (
	"bytes"
	"io"
	"os"
)

func CaptureStdErr(f func()) string {
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	f()

	_ = w.Close()
	os.Stderr = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
