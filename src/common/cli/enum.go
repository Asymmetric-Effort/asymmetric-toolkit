package cli

func (o *Configuration) Enum(name string, defaultValue string, enumeratedValues ...string) ValidationFunc {
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