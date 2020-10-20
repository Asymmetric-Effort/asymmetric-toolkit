package logLevel

import (
	"fmt"
)

func (o *LogLevel) Get() (v LogLevel) {
	switch *o {
	case Critical, Error, Warning, Info, Debug:
		v = *o
		break
	default:
		panic(fmt.Sprintf("Programming error: invalid Log level: %d", *o))
	}
	return
}
