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
	lineFmt   string = "\t%s%s [Default(%s): %s]\n"
)

func (o *Specification) calcArgColumnWidth() int {
	max:=0
	for flag, _ := range o.Argument {
		flagLength:=len(flag)
		if flagLength > max {
			max = flagLength
		}
	}
	return max
}

func (o *Specification) ShowHelp(arg *string) (err error, val *Argument) {
	/*
		Calculate and show the usage message (all help messages).
	*/
	errors.Assert(arg != nil, "Expected non nil argument.")
	ArgColumnWidth:=o.calcArgColumnWidth()
	output:=""
	output = fmt.Sprintf(bannerFmt, o.ProgramName, o.Version, o.Copyright, o.Description)
	spacer:=func(n int)string{
		s:=""
		for i:=0;i<n;i++{
			s+=" "
		}
		return s
	}
	for flag, arg := range o.Argument {
		flagStr := func() (f string) {
			flagLength:=len(flag)
			spacerLen:=(ArgColumnWidth-flagLength)+5
			if flagLength == 1 {
				f = fmt.Sprintf("-%s%s", flag,spacer(spacerLen-1))
			} else {
				f = fmt.Sprintf("--%s%s", flag,spacer(spacerLen-2))
			}
			return f
		}()
		output += fmt.Sprintf(lineFmt, flagStr, arg.Help, arg.Type.String(), arg.Default)
	}
	fmt.Printf(output)
	return nil, nil
}
