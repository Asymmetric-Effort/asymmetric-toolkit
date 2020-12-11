package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
)

func (o *File) MoveEof() int64 {
	pos, err := o.Handle.Seek(0, 2) //EOF
	errors.Assert(err == nil, fmt.Sprintf("%v", err))
	return pos
}