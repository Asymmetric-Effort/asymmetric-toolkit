package types

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"regexp"
)

type FilterPattern struct {
	re *regexp.Regexp
}

func (o *FilterPattern) Set(s string) {
	re, err := regexp.Compile(s)
	errors.Assert(err == nil, fmt.Sprintf("Encountered an error when compiling --pattern regex.  Error:%v", err))
	o.re = re
}

func (o *FilterPattern) Match(s string) bool {
	if o.re == nil {
		return false
	}
	return o.re.MatchString(s)
}

func (o *FilterPattern) String() string {
	if o.re == nil {
		return ""
	}
	return o.re.String()
}
