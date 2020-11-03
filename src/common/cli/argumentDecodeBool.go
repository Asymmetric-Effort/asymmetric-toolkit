package cli

import "strings"

func (o *Argument) DecodeBool() bool {
	if o.Type == Boolean {
		switch strings.TrimSpace(strings.ToLower(o.Value)) {
		case "true":
			return true
		case "false":
			return false
		default:
			panic("Unexpected value encountered when returning a boolean.  Expected true/false.")
		}
	} else {
		panic("Type mismatch.  DecodeBool() called for non-boolean Argument type.")
	}
}