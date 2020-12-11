package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
)

func (o *File) WriteHeader(header *Header) {
	if header != nil {
		//If input is nil pointer, we will not update our struct.
		//which means we will write whatever state is in memory to disk.
		o.Header = header
	}
	_, o.Err = o.Header.Write(o.Handle)
	errors.Assert(o.Err == nil, fmt.Sprintf("%v", o.Err))
}
