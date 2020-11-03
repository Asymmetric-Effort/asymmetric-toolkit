package cli

import "strconv"

func (o *Argument) DecodeInt() int {
	if o.Type == Integer {
		i, err := strconv.Atoi(o.Value)
		if err != nil {
			panic(err)
		}
		return i
	} else {
		panic("Type mismatch.  DecodeInt() called for non-integer Argument type.")
	}
}