package cli

/*
	Extractor function Argument::Boolean() will convert the internal (string) state to an boolean
	and return the same boolean value to the caller.
*/

import "strconv"

func (o *Argument) Boolean() bool {
	if o.Type == Boolean {
		return func() bool {
			i, err := strconv.ParseBool(o.Value)
			if err != nil {
				panic(err)
			}
			return i
		}()
	} else {
		panic("Type mismatch.  Attempted to extract float from non-float Argument.")
	}
}
