package types

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"regexp"
	"strings"
)

type DomainName string

func (o *DomainName) Set(s string) {
	re:=regexp.MustCompile(`^[A-Za-z0-9-_\\.]+$`)
	t:=strings.TrimSpace(s)
	errors.Assert(re.MatchString(t), "Domain name failed validation.")
	*o = DomainName(t)
}

func (o *DomainName) Get() string {
	return string(*o)
}
