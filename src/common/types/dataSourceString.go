package types

import (
	"fmt"
)

func (o *DataSource) String() string {
	switch *o {
	case Random:
		return "Random"
	case Sequence:
		return "Sequence"
	case Dictionary:
		return "Dictionary"
	default:
		panic(fmt.Sprintf("Cannot convert (%d) to string.  Unrecognized.", *o))
	}
}
