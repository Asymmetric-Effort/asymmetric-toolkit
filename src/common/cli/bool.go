package cli

import (
	"strings"
)

func (o *Configuration) Bool(name string, defaultValue bool) ValidationFunc {
	return func(i *string) bool {
		switch strings.ToLower(*i) {
		case "true":
			o.args[name] = true
		case "false":
			o.args[name] = false
		case "": //No input.  Go with default.
			o.args[name] = defaultValue
		default:
			return false //Invalid input.
		}
		return true
	}
}