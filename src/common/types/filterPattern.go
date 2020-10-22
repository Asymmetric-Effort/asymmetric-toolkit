package types

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"regexp"
)

type FilterPattern struct {
	re *regexp.Regexp
}

const (
	DefaultFilterPattern string = `.+`
)

func (o *FilterPattern) Set(s string) {
	re, err := regexp.Compile(s)
	errors.Assert(err == nil, fmt.Sprintf("Encountered an error when compiling --pattern regex.  Error:%v", err))
	o.re = re
}

func (o *FilterPattern) Match(s string) bool {
	if o.IsNil() {
		return false
	}
	return o.re.MatchString(s)
}

func (o *FilterPattern) String() string {
	if o.IsNil() {
		return ""
	}
	return o.re.String()
}

func (o *FilterPattern) Get() string {
	return o.String()
}

func (o *FilterPattern) IsNil() bool {
	return o.re == nil
}
