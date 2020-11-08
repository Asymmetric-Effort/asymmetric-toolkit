package cli

/*
	ShowUsage is a ParserFunction, following the same pattern.  However, this parser function will not actually
	parse any argument or create any value.  It simply calculates then displays program usage information (flags, etc).
*/

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
)

func (o *Specification) ShowUsage(arg *string) (err error, val *Argument) {
	/*
		Calculate and show the usage message (all help messages).
	*/
	errors.Assert(arg != nil, "Expected non nil argument.")
	fmt.Printf("\n" +
		"%s (%s)\n\n" +
		"%s\n\n" +
		"Usage:\n",o.ProgramName,o.Version,o.Copyright)
	for flag, arg := range o.Argument {
		fmt.Printf("\t%s\t\t%s\t[Default: %s]\n", flag, arg.Help, arg.Default)
	}
	return nil, nil
}
