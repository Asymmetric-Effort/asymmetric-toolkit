package cli

import "fmt"

/*
	CommandLine::SetDefaults() will iterate through the specification and create Argument objects for each
	known flag and set the default values in the specification for each flag.

	Note: This will also initialize each argument
*/
func (o *CommandLine) SetDefaults(spec *Specification) (err error) {
	//
	// Iterate through the specification
	//
	for flag, flagSpec := range spec.Argument {
		//
		// Call the parser function for each argument and pass the default value.
		// This ensures the default value is validated just like a user-provided value,
		// then process any error.
		//
		err, o.Arguments[flagSpec.FlagId] = flagSpec.Parse(&flagSpec.Default)
		//
		// Stop processing on error
		//
		return fmt.Errorf("flag:%s, error: %v", flag, err)
	}
	//
	// return result (no error expected here).
	//
	return
}
