package types

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDomainName_SetHappy(t *testing.T) {
	var d DomainName
	d.Set("google.com")
	d.Set("mySub.domain.tld")
}
func TestDomainName_SetBad(t *testing.T) {
	var d DomainName
	defer func(){recover()}()
	d.Set("Bad$ub.D0Ma!n.tld")
}
func TestDomainName_SetEmpty(t *testing.T) {
	var d DomainName
	defer func(){recover()}()
	d.Set("")
}

func TestDomainName_Get(t *testing.T) {
	var d DomainName
	const testDomain="subdomain.domain.tld"
	d.Set(testDomain)
	errors.Assert(d.Get() == testDomain, "DomainName::Get() failed")
}