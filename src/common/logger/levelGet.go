package logger

/*
	Level::Get() returns the numeric Level for the current state and validates the same.
*/
import (
	"fmt"
)

func (o *Level) Get() (v Level) {
	switch *o {
	case Critical, Error, Warning, Info, Debug:
		v = *o
		break
	default:
		panic(fmt.Sprintf("Programming error: invalid Log level: %d", *o))
	}
	return
}
