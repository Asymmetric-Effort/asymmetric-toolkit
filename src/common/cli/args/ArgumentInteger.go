package cliargs

import (
	"fmt"
	"strconv"
)

/*
	The ArgumentInteger struct implements the Arguments interface{}
	and decodes the string cli parameter into the integer value it
	represents.
 */

type ArgumentInteger struct {
	data int
}

func (o *ArgumentInteger) Get() int {
	return o.data
}

func (o *ArgumentInteger) Set(n int) {
	o.data = n
}


func (o *ArgumentInteger) SetString(s string, low int, high int, defaultValue int) error {
	if s==""{
		if defaultValue < low || defaultValue > high {
			panic("defaultValue out of range.")
		}
		o.data = defaultValue
		return nil
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("%v",err)
	}
	if n >= low && n <= high {
		o.data = n
	}
	return nil
}
