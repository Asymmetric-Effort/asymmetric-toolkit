package cli
/*
	To implement this module, we define a Specification object.  This is passed to the CommandLine::Parse() method
	by reference to define how the CommandLine processor will interpret existing commandline parameters to configure
	the internal state of the system.
*/

type CommandLine struct {
	Arguments map[ArgumentFlag]*Argument
}
