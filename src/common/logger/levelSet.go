package logger

/*
	Level::Set() provides a method for validating and setting a Level numeric value.
*/
import (
	"fmt"
)

func (o *Level) Set(v Level) {
	switch v {
	case Critical, Error, Warning, Info, Debug:
		break
	default:
		panic(fmt.Sprintf("Programming error: invalid Log level: %d", v))
	}
	*o = v
}
