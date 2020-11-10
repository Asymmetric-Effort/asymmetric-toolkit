package cli
/*
	The ArgumentDescriptor exists to represent the commandline argument inside the specification provided
	by the programmer.  This descriptor describes how the argument will be processed into its final,
	usable state.  This is used only by the program to instruct the commandline processor on how to
	process the commandline arguments, and it is expected to be freed from memory once commandline
	processing is done.
 */

type ArgumentDescriptor struct {
	FlagId	ArgumentFlag   // A numeric identifier for the flag used within the standard flags.
	Type    ArgumentType   // An expected data type.
	Default string         // The default value
	Help    string         // Help text
	Parse   ParserFunction // Argument-specific parser which will set the internal string state or return an error.
	Expects NextExpected
}
