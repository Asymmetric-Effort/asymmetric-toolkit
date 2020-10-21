package writer

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
)


func (o *Writer) Close() {
	if o.file != nil {
		err := o.file.Close()
		errors.Assert(err == nil, fmt.Sprintf("Failed to close file.  Errorf:%v", err))
	}
}