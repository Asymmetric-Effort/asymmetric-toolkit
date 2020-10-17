package reader

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
)

func (o *Reader) Close() {
	if o.file != nil {
		err := o.file.Close()
		errors.Assert(err == nil, fmt.Sprintf("Failed to close file.  Error:%v", err))
	}
}
