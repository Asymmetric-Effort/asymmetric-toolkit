package cli

func (o *Argument) DecodeString() string {
	if o.Type == String {
		return o.Value
	} else {
		panic("Type mismatch.  DecodeString() called for non-string Argument type.")
	}
}