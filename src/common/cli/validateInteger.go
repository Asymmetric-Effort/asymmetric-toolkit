package cli

import cliargs "asymmetric-effort/asymmetric-toolkit/src/common/cli/args"

func (o *CommandLine) ValidateInteger(name string, defaultValue int, low int, high int) ValidationFunc {
	if low > high {
		panic("Internal Error: Expect low < high in valid.IntRange()")
	}
	/*
		Create and return a validation function
	*/
	return func(i *string) bool {
		var intArg cliargs.ArgumentInteger
		err := intArg.SetString(*i, low, high, defaultValue)
		if err != nil {
			panic(err)
		}
		o.args[name] = intArg.Get()
		return true
	}
}
