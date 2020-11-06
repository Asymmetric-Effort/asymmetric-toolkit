package cli

/*
	Extractor function Argument::Integer() will convert the internal (string) state to an integer
	and return the same integer value to the caller.
*/

import "strconv"

func (o *Argument) Integer() int {
	if o.Type == Integer {
		return func() int {
			i, err := strconv.Atoi(o.Value)
			if err != nil {
				panic(err)
			}
			return i
		}()
	} else {
		panic("Type mismatch.  Attempted to extract integer from non-integer Argument.")
	}
}