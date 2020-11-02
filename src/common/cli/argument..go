package cli
/*
	The Argument represents the processed commandline argument
	still in string format but validated.  The program can use
	this struct to extract integers, booleans, floats, etc. into
	the program's internal state.
 */

type Argument struct {
	Type  ArgumentType
	Value string		// This is the parsed, validated but not decoded (i.e. string) value parsed from the cli.
}
