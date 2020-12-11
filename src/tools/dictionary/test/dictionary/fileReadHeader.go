package dictionary
/*
	dictionary.File::ReadHeader() reads the dictionary file header.
 */
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
)

func (o *File) ReadHeader() {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if o.Header == nil {
		o.Header = &Header{}
	}
	o.Err = o.Header.Read(o.Handle)
	errors.Assert(o.Err == nil, fmt.Sprintf("%v", o.Err))
}
