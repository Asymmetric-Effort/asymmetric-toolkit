package cli

/*
	Extractor function Argument::Float() will convert the internal (string) state to an Float64
	and return the same Float64 value to the caller.
*/

import "strconv"

func (o *Argument) Float() float64 {
	if o == nil {
		panic("Nil Float argument pointer")
	}
	if o.Type == Float {
		return func() float64 {
			i, err := strconv.ParseFloat(o.Value, 64)
			if err != nil {
				panic(err)
			}
			return i
		}()
	} else {
		panic("Type mismatch.  Attempted to extract float from non-float Argument.")
	}
}
