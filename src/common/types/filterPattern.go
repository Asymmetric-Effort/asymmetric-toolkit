package types

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"regexp"
)

type FilterPattern struct {
	Re *regexp.Regexp
}

func (o *FilterPattern) Set(s string) {
	re, err := regexp.Compile(s)
	errors.Assert(err == nil, fmt.Sprintf("Encountered an error when compiling --pattern regex.  Errorf:%v", err))
	o.Re = re
}

func (o *FilterPattern) Match(s string) bool {
	if o.Re == nil {
		return false
	}
	return o.Re.MatchString(s)
}

func (o *FilterPattern) String() string {
	if o.Re == nil {
		return ""
	}
	return o.Re.String()
}
