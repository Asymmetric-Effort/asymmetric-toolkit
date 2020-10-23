package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func (o *Configuration) Setup(spec *Specification) (err error) {
	var currentFlag string
	var count int         //count the flags for later comparison.
	var next = ExpectFlag // We always expect the first arg to be a --flag.

	o.spec = logger.CommandLineSpecification(o)           // Get the logger cli specification.
	MergeMaps(o.spec, source.CommandLineSpecification(0)) // Adding source CLI.
	MergeMaps(o.spec, spec)                               // This is the design of our cli expectation.

	o.args = make(Arguments) // Initialize the arguments object for use before parsing.
	/*
		Parse the arguments
	*/
	for _, arg := range os.Args[1:] {
		switch next {
		case expectFlag:
			if isFlag(&arg) {
				currentFlag = getFlag(&arg)
			} else {
				return fmt.Errorf("A flag was expected but not found.  " +
					"Please use -h or --help for usage information.")
			}
			if o.specificationExpectsFlag(&currentFlag) {
				if (*o.spec)[currentFlag].ExpectValue {
					next = expectValue
					count++
					continue
				} //else we're gonna grab another flag.
			} else {
				return fmt.Errorf("This tool expects a certain number of flags and parameters.  "+
					"But the given flag (%s) was not expected.\n", currentFlag)
			}
		case ExpectValue:
			if (*o.spec)[currentFlag].Validation(&arg) { // Set, validate flag value and move on.
				next = expectFlag
				continue
			} else {
				return fmt.Errorf("Unexpected value for %s flag.  Encountered '%s'.  "+
					"Please use -h or --help for usage information.\n", currentFlag, arg)
			}
		default:
			return fmt.Errorf("internal error parsing command line options")
		}
	}
	if count != o.countRequiredArgs() {
		return fmt.Errorf("One or more missing but expected flags.  " +
			"Please use -h or --help for usage information.")
	}
	/*
		Perform a second pass to ensure all expectations are met.  For some args this means taking action.
	*/
	if _, ok := o.args[FlagHelp]; ok {
		o.ShowHelp()
	}
	if _, ok := o.args[FlagVersion]; ok {
		o.ShowVersion()
	}
	/*
		We are done. Return whatever error state we are in at this point.  Most likely we are nil.
	*/
	return err
}

func getFlag(s *string) string {
	re := regexp.MustCompile("^--.+$")
	if re.MatchString(*s) {
		return strings.ToLower((*s)[2:])
	} else {
		panic("Expected flag but encountered non flag arg.")
	}
}

func isFlag(s *string) bool {
	re := regexp.MustCompile("^--.+$")
	return re.MatchString(*s)
}
