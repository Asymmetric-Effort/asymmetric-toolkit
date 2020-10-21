package types

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strconv"
)

type PositiveInteger int

func (o *PositiveInteger) Set(s string) {
	i, e := strconv.Atoi(s)
	errors.Assert(e == nil, fmt.Sprintf("Expected positive integer. Error:%v", e))
	errors.Assert(i > 0, "Numeric Value must be greater than zero(0).")
	*o = PositiveInteger(i)
}

func (o *PositiveInteger) Get() int{
	return int(*o)
}