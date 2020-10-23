package cli

func (o *CommandLine) Enum(name string, defaultValue string, enumeratedValues EnumeratedValues) ValidationFunc {
/*	return func(i *string) ValidationFunc {
		return func(i *string) bool {
			var retval string
			if *i == "" {
				retval = def
			}
			return
		}
	}*/
	return Noop
}