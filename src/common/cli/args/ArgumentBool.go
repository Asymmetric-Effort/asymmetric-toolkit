package cliargs

import "strings"

type ArgumentBool struct {
	data bool
}

func (o *ArgumentBool) Get() bool {
	return o.data
}

func (o *ArgumentBool) Set(n bool) {
	o.data = n
}

func (o *ArgumentBool) SetString(s string, defaultValue bool) {
	switch strings.ToLower(s) {
	case "":
		o.data = defaultValue
	case "true":
		o.data = true
	case "false":
		o.data = false
	default:
		panic("ArgumentBool::SetString encountered non-boolean input.")
	}
}
