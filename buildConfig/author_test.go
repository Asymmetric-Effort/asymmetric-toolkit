package buildconfig_test

import (
	buildconfig "asymmetric-effort/asymmetric-toolkit/buildConfig"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestAuthor(t *testing.T) {
	errors.Assert(buildconfig.Author == "Sam Caldwell", "expected author mismatch")
}
