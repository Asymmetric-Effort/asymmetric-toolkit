package cli

import (
	cliargs "asymmetric-effort/asymmetric-toolkit/src/common/cli/args"
	"regexp"
)

func (o *CommandLine) ValidateString(name string, defaultValue string, pattern string) ValidationFunc {
	re := regexp.MustCompile(pattern)
	/*
		Create and return a validation function
	*/
	return func(i *string) (result bool) {
		if re.MatchString(*i) {
			var strArg cliargs.ArgumentString
			strArg.SetString(*i, defaultValue)
			o.args[name] = strArg.Get()
			return true
		}
		return false
	}
}
