package cli

import (
	"regexp"
	"strings"
)

/*
	stripPrefix() is used by CommandLine::Parse() to strip the leading - or -- from a given flag
	and return the remaining string.  Thus, "-h" becomes "h" and "--help" becomes "help".

	Valid flag patterns include:
		-(char)
		--(>=2chars)
*/
func stripPrefix(arg *string) (isFlag bool, flag string) {
	const regExFlagPattern = `^([-]{1}.{1}|[-]{2}.{64})$`
	re := regexp.MustCompile(regExFlagPattern)
	flag = strings.TrimLeft(*arg, "-")
	isFlag = re.MatchString(*arg) && (len(flag) > 0)
	return
}
