package level

import (
	"fmt"
)

func (o *Level) Set(v Level) {
	switch v {
	case Critical, Error, Warning, Info, Debug:
		break
	default:
		panic(fmt.Sprintf("Programming error: invalid Log level: %d", v))
		*o = Debug
		return
	}
	*o = v
}
