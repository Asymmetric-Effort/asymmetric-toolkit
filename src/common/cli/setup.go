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
	/*
		This is where we aggregate (left join) our commandline specifications
		into a single final view.  Each subsequent layer can override the previous.
	*/
	o.spec = CommandLineSpecification(o)
	MergeMaps(o.spec, logger.CommandLineSpecification(o))
	MergeMaps(o.spec, source.CommandLineSpecification(o))
	MergeMaps(o.spec, spec)

	// Initialize the arguments object for use before parsing.
	o.args = make(Arguments)
	/*
		Parse the arguments
	*/
	for _, arg := range os.Args[1:] {
		switch next {
		case expectFlag:
			// Is the argument a flag?....
			isFlagRe:=regexp.MustCompile("^--.+$")
			if !isFlagRe.MatchString(arg) {
				return fmt.Errorf("A flag was expected but not found.  " +
					"Please use -h or --help for usage information.")
			}
			currentFlag = func() string {
				// Extract the --flag
				re := regexp.MustCompile("^--.+$")
				if re.MatchString(arg) {
					return strings.ToLower((arg)[2:])
				} else {
					panic("Expected flag but encountered non flag arg.")
				}
			}()
			count++
			if !o.specificationExpectsFlag(&currentFlag) {
				return fmt.Errorf("This tool expects a certain number of flags and parameters.  "+
					"But the given flag (%s) was not expected.\n", currentFlag)
			}
			if (*o.spec)[currentFlag].ExpectValue {
				next = expectValue
			} // else we're gonna grab another flag.
		case ExpectValue:
			if !(*o.spec)[currentFlag].Validation(&arg) { // Set, validate flag value and move on.
				return fmt.Errorf("Unexpected value for %s flag.  Encountered '%s'.  "+
					"Please use -h or --help for usage information.\n", currentFlag, arg)
			}
			next = expectFlag
			continue

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
