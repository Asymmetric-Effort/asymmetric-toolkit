package types

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"github.com/graph.baby/src/common/file"
	"os"
	"strings"
)

type FilePath string

func (o *FilePath) Set(s string) {
	fp := strings.TrimSpace(s)
	errors.Assert(fp != "", "Expected non-empty string for file/path.")
	*o = FilePath(fp)
}

func (o *FilePath) Get() string {
	return string(*o)
}

func (o *FilePath) Exists() bool {
	return file.FileExists(string(*o))
}

func (o *FilePath) OpenRead() *os.File {
	fp, err := os.Open(string(*o))
	if err != nil {
		errors.Fatal(1, fmt.Sprintf("File open error.  %v", err))
	}
	return fp
}

func (o *FilePath) OpenWrite() *os.File {
	panic("not implemented")
}
