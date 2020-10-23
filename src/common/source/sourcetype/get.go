package sourcetype

import (
	"fmt"
)

func (o *DataSource) Get() DataSource {
	switch *o {
	case NotSet, Random, Sequence, Dictionary:
		return *o
	default:
		panic(fmt.Sprintf("Invalid DataSource value (%d) in Get().  "+
			"Be sure to use .Set() to set values.   Direct assignment of non enumerated values will cause "+
			"errors.", *o))
	}
}
