package DictionaryReader_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"os"
	"testing"
)

func TestReaderClose(t *testing.T){
	var r Reader
	var err error
	r.file, err = os.Open(os.Args[0])
	errors.Assert(err == nil, fmt.Sprintf("%v", err))

}