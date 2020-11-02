package cli

import (
	"fmt"
	"os"
)

func (o *CommandLine) Parse(spec *Specification) (exit bool, err error) {
	/*
		Parse the existing command line arguments, perform validation and store the
		values into the internal state.

		Input: *Specification (describe how the argument parser will work.
		Output:
			- Should the caller exit or continue? (default: continue.
	*/
	expected := expectFlag
	lastFlag := noFlag

	spec.AddUsage()

	for _, arg := range os.Args[1:] {
		switch expected {
		case expectFlag:
			flagDetected := false
			for flag, flagSpec := range spec.Argument {
				switch arg {
				case flag:
					lastFlag = flagHelp
					expected = expectEnd
					flagSpec.Parse()
				}
			}
			if flagDetected { // determine if flag is detected in specification.
				continue
			} else {
				return true, fmt.Errorf("missing or unrecognized argument: %s", arg)
			}
		case expectValue:
			switch lastFlag {

			default:
				return true, fmt.Errorf("expected value for unknown lastFlag: %d", lastFlag)
			}
		case expectEnd:
			//No other arguments will be processed.  We expect to exit.
			return true, nil
		default:
			return true, fmt.Errorf("unexpected argument parser state. Expected: %d", expected)
		}
	}
	return false, nil
}
