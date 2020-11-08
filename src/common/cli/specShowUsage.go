package cli

/*
	ShowUsage is a ParserFunction, following the same pattern.  However, this parser function will not actually
	parse any argument or create any value.  It simply calculates then displays program usage information (flags, etc).
*/

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
)

const (
	bannerFmt string = "\n%s (%s) %s\n\nUsage:\n"
	lineFmt   string = "\t%s\t\t%s [Default: %s]\n"
)

func (o *Specification) ShowUsage(arg *string) (err error, val *Argument) {
	/*
		Calculate and show the usage message (all help messages).
	*/
	errors.Assert(arg != nil, "Expected non nil argument.")
	output := ""
	output = fmt.Sprintf(bannerFmt, o.ProgramName, o.Version, o.Copyright)
	for flag, arg := range o.Argument {
		output += fmt.Sprintf(lineFmt, flag, arg.Help, arg.Default)
	}
	fmt.Printf(output)
	return nil, nil
}
