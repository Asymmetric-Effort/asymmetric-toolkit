package cli

/*
	CommandLine::Parse() is the top-level commandline argument parser which takes a given Specification object
	and applies it to the command line arguments in os.Args[1:] to produce a set of Argument objects the program
	can use to configure its internal state.
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"fmt"
)

func (o *CommandLine) Parse(spec *Specification, args *[]string) (exit bool, err error) {
	/*
		Parse the existing command line arguments, perform validation and store the
		values into the internal state.

		Input: *Specification (describe how the argument parser will work.
		Output:
			- Should the caller exit or continue? (default: continue.
	*/
	var expected = ExpectFlag // NextExpected
	var lastFlag *ArgumentDescriptor = nil

	if len(*args) == 0 { // No argument?  Exit.
		return true, fmt.Errorf("%s",ErrMsgMissingArguments)
	}

	spec.AddHelp()    // If our help flags (-h and --help) are not set, we will add them here.
	spec.AddVersion() // If our version flags (-v and --version) are not set, we will add them here.
	spec.AddDebug()   // If our debug flag (--debug) is not set, we will add it here.
	spec.AddForce()   // If our force flag (--force) is not set, we will add it here.
	spec.AddLogLevel(logger.Info)
	spec.AddLogDestination(logger.Stdout)

	err = spec.EnsureUniqueFlagId() //Scan the specification and ensure we have unique flagIDs.
	if err != nil {
		panic(fmt.Sprintf("FlagId uniqueness error: %v", err))
	}
	//
	// Set the default values for our specification.
	//
	if err := o.SetDefaults(spec); err != nil {
		return true, fmt.Errorf("error applying default values in commandline processor. %v", err)
	}

	for _, currentArgument := range *args {
		//
		// Iterate through all arguments on the commandline.  We expect either
		// a long flag (--flag), a short flag (-f) or a contiguous string value.
		//
		switch expected {
		case ExpectFlag:
			//
			// If we expect a flag (--long or -l), we will parse the flag and
			// determine its specifications to process the next argument.
			//
			isFlag, arg := stripPrefix(&currentArgument)
			if !isFlag {
				//
				// What we encountered is not a flag, but some other identifier.
				// This is an error.  Terminate execution.
				//
				return true, fmt.Errorf("expected flag but encountered "+
					"non-flag argument (%s)", currentArgument)
			}
			if knownSpec, ok := spec.Argument[arg]; ok {
				//
				// For the current argument, does it exist in our specification's argument map?
				// If it does exist, we have a detected flag we can process...
				//
				lastFlag = &knownSpec        // Save the current ArgumentDescriptor for the flag.
				expected = knownSpec.Expects // Route the next argument as expected in the spec.
				//
				if expected == ExpectNone {
					//
					// If we expect nothing (e.g. --help, --debug, --verbose) where the flag is the value
					// in and of itself, we simply process it now rather than wait another cycle (which would process
					// the next argument).
					//
					if err, o.Arguments[lastFlag.FlagId] = lastFlag.Parse(&currentArgument); err == nil {
						if o.Arguments[lastFlag.FlagId] == nil {
							// If our flag Parse() function returned nil in this case
							// we would know the Parse() function is terminal and we should exit without error,
							// as in the --help or --version use cases.
							return true, nil
						}
						// We have hit a flag, which is present (e.g. --debug or --force) and we need to set true.
						o.Arguments[lastFlag.FlagId].Value = "true"
						expected = ExpectFlag // Reset and get another flag.
					} else {
						//
						//We encountered an error on our last flag (expectNone parser function).
						//
						return true, err
					}
				}
			} else {
				//
				// Terminate with an error.  We failed to find the current argument as an
				// argument in our current specification.
				//
				return true, fmt.Errorf("missing or unrecognized argument: %s", arg)
			}
			continue
		case ExpectValue:
			//
			// If we expect a value, we presume a flag was detected first to which we can map the
			// parsed value currently in arg.
			//
			if lastFlag == nil {
				//
				// If we have no argument (arg) flag, we return an error because we cannot map a value
				// without a flag to represent what the value actually means.
				//
				fmt.Println("lastFlag is nil")
				return true, fmt.Errorf("a value was found with no flag: %s", currentArgument)
			}
			//
			// Our lastFlag (ArgumentDescriptor) defines how we should proceed to parse, validate and
			// handle the value currently in arg.  If successful, the CommandLine (o) structure will
			// contain the parsed (string) value in the Argument map.
			//
			if lastFlag.Parse == nil {
				//
				// We have a nil pointer reference for our parser.  To avoid a crash we simply
				// continue at this point and assume a noop.  No value will be stored for this
				// argument and move on to the next flag.
				//
				expected = ExpectFlag
			} else {
				var err error
				//
				// If a Parse() function is provided for validation, we will execute the same
				// to parse and validate our data specific to its context. If no error occurs,
				// We expect the Parse() function will return an Argument object containing the
				// processed value and any additional features.
				//
				if err, o.Arguments[lastFlag.FlagId] = lastFlag.Parse(&currentArgument); err == nil {
					//
					// We have properly processed the cli argument value and we will now reset
					// to expect a flag next to start the process over again.
					//
					expected = ExpectFlag
				} else {
					//
					// For each ArgumentDescriptor, a Parse() function is provided.  This function
					// will perform any context-specific validation and store the value as an Argument
					// object in CommandLine (o).
					//
					return true, err
				}
				if o.Arguments[lastFlag.FlagId] == nil {
					//
					// After parsing the current argument value from the commandline, we perform a nil
					// pointer check to avoid application crashes.  Parse() functions cannot return nil
					// Arguments without returning an error.  A nil Argument implies that the argument
					// failed to parse/validate.
					//
					return true, fmt.Errorf("our flag parse() function cannot return a nil "+
						"Argument pointer (%d)", lastFlag.FlagId)
				}
			}
			continue
		case ExpectEnd:
			//
			// No other arguments will be processed.  We expect to exit the program.
			//
			return true, nil
			//
		default:
			//
			// We have reached an internal programming error.  We expected something as a NextExpected
			// value which is not known to Commandline::Parse().  At this point, we are forced to abandon
			// all hope and terminate execution.
			//
			return true, fmt.Errorf("unexpected argument parser state. Expected: %d", expected)
			//
		}
	}
	//
	// We have parsed all arguments and none of the arguments encountered errors or instructed our program
	// to terminate.  At this point we will set exit = false and return a nil error object and allow the
	// program to continue execution as designed.
	//
	return false, nil
}
