package cli

import "strconv"

func (o *Argument) DecodeFloat32() float32 {
	if o.Type == Float {
		n, err := strconv.ParseFloat(o.Value, 32)
		if err != nil {
			panic(err)
		} else {
			return float32(n)
		}
	} else {
		panic("Type mismatch.  DecodeFloat32() called for non-float Argument type.")
	}
}
