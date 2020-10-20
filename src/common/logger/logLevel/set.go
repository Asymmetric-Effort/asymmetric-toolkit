package logLevel

import (
	"fmt"
)

func (o *LogLevel) Set(v LogLevel) {
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
