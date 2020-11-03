package cli

import "strconv"

func (o *Argument) DecodeFloat64() float64 {
	if o.Type == Float {
		n, err := strconv.ParseFloat(o.Value, 64)
		if err != nil {
			panic(err)
		} else {
			return n
		}
	} else {
		panic("Type mismatch.  DecodeFloat64() called for non-float Argument type.")
	}
}
