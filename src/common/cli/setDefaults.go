package cli

import (
	"fmt"
	"strconv"
)

/*
	CommandLine::SetDefaults() will iterate through the specification and create Argument objects for each
	known flag and set the default values in the specification for each flag.

	Note: This will also initialize each argument
*/
func (o *CommandLine) SetDefaults(spec *Specification) (err error) {
	//
	// Iterate through the specification
	//
	o.Arguments = make(map[ArgumentFlag]*Argument)
	delete(o.Arguments, noFlag)
	delete(spec.Argument, "")
	for flag, flagSpec := range spec.Argument {
		//fmt.Println("flag:", flag, " flagSpec:", flagSpec)
		//
		// Call the parser function for each argument and pass the default value.
		// This ensures the default value is validated just like a user-provided value,
		// then process any error.
		//
		if (flagSpec.Expects == ExpectNone) || (flagSpec.Expects == ExpectEnd) {
			//fmt.Println("\tExpects:", flagSpec.Expects, "(none|end) Default:", flagSpec.Default)
			_, err := strconv.ParseBool(flagSpec.Default)
			if err != nil {
				panic(fmt.Sprintf("(%d) Expected boolean default value.  Encountered non-boolean.",
					flagSpec.FlagId))
			}
			o.Arguments[flagSpec.FlagId] = &Argument{
				Boolean,
				flagSpec.Default,
			}
		} else {
			err, o.Arguments[flagSpec.FlagId] = flagSpec.Parse(&flagSpec.Default)
			//
			// Stop processing on error
			//
			if o.Arguments[flagSpec.FlagId] == nil {
				panic(fmt.Sprintf("flag:'%s', error: missing or invalid default value", flag))
			}
			if err != nil {
				err = fmt.Errorf("flag:%s, error: %v", flag, err)
			}
		}
	}

	//
	// return result (no error expected here).
	//
	return err
}
