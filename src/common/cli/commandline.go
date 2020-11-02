package cli
/*
	To implement this module, we define a Specification object.  This is passed to the CommandLine::Parse() method
	by reference to define how the CommandLine processor will interpret existing commandline parameters to configure
	the internal state of the system.
*/

type CommandLine struct {
	Arguments map[string]Argument
	/*
		Decoder functions are setup by the Parse() method which gets them from the Specification object.
		These functions are used to craft argument-specific features to process specific arguments by type
		and return a validated, sanitized state.
	*/
	DecodeString map[string]func() string
	DecodeInt    map[string]func() int
	DecodeBool   map[string]func() bool
	DecodeFloat  map[string]func() float64
	DecodeEnum   map[string]func() string
}
