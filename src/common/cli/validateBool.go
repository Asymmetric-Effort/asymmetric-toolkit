package cli

import cliargs "asymmetric-effort/asymmetric-toolkit/src/common/cli/args"

func (o *CommandLine) ValidateBool(name string, defaultValue bool) ValidationFunc {
	/*
		Create and return a validation function
	*/
	return func(i *string) bool {
		var boolArg cliargs.ArgumentBool
		boolArg.SetString(*i, defaultValue)
		return true
	}

}
