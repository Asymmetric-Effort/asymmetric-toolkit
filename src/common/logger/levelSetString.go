package logger

/*
	Level::SetString() provides a method for validating and setting a Level numeric value given the string equivalent.
*/
import (
	"fmt"
	"strings"
)

func (o *Level) SetString(v string) {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "critical": *o=Critical
	case "error": *o=Error
	case "warning": *o=Warning
	case "info": *o=Info
	case "debug": *o=Debug
	default:
		panic(fmt.Sprintf("Invalid log level (%s)",v))
	}
}
