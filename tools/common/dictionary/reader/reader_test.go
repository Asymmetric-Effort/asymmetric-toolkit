package DictionaryReader

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	var r Reader
	var err error
	errors.Assert(r.file == nil, fmt.Sprintf("Expected nill file handle"))
	r.file, err = os.Open(os.Args[0])
	errors.Assert(err == nil, fmt.Sprintf("%v", err))
	errors.Assert(r.file != nil, fmt.Sprintf("Expected nill file handle"))
}