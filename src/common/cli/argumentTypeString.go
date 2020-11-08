package cli

/*
	The ArgumentType::String method converts a numeric ArgumentType into its string equivalent.
*/
import (
	"fmt"
	"strings"
)

func (o *ArgumentType) String() string {
	//
	// Split the comma-delimited string into an array
	//
	argTypes := strings.Split(argumentTypes, comma)
	//
	// Bounds check our ArgumentType (numeric value) against the known list of strings.
	//
	if (int(*o) >= 0) && (int(*o) < len(argTypes)) {
		//
		// Return the string corresponding to the numeric value (cast as int)
		//
		return argTypes[int(*o)]
	} else {
		//
		// If bounds check fails, look up will fail.  Panic.
		//
		panic(fmt.Sprintf("Cannot convert ArgumentType (%d) into String. %s", *o, argumentTypes))
	}
}
