package buildconfig_test

import (
	buildconfig "asymmetric-effort/asymmetric-toolkit/buildConfig"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestAuthorEmail(t *testing.T) {
	errors.Assert(buildconfig.AuthorEmail == "asymmetric-toolkit@samcaldwell.net",
		"expected author email mismatch")
}
