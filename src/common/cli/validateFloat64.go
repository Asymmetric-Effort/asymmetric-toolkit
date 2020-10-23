package cli

func (o *CommandLine) ValidateFloat64(name string, defaultValue float64) ValidationFunc {
	/*
		Create and return a validation function
	*/
	return func(i *string) bool {
		panic("ValidateFloat64 not implemented yet.")
	}
}
