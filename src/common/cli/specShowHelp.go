package cli

/*
	ShowHelp is a ParserFunction, following the same pattern.  However, this parser function will not actually
	parse any argument or create any value.  It simply calculates then displays program usage information (flags, etc).
*/

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
)

const (
	bannerFmt string = "\n%s (%s) %s\n\n%s\n\nUsage:\n"
	lineFmt   string = "\t%s\t\t%s [Default(%s): %s]\n"
)

func (o *Specification) ShowHelp(arg *string) (err error, val *Argument) {
	/*
		Calculate and show the usage message (all help messages).
	*/
	errors.Assert(arg != nil, "Expected non nil argument.")
	output := ""
	output = fmt.Sprintf(bannerFmt, o.ProgramName, o.Version, o.Copyright, o.Description)
	for flag, arg := range o.Argument {
		flagStr := func() (f string) {
			if len(flag) == 1 {
				f = fmt.Sprintf("-%s", flag)
			} else {
				f = fmt.Sprintf("--%s", flag)
			}
			return f
		}()
		output += fmt.Sprintf(lineFmt, flagStr, arg.Help, arg.Type.String(), arg.Default)
	}
	fmt.Printf(output)
	return nil, nil
}
