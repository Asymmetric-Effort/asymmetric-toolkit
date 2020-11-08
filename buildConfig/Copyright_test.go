package buildconfig_test

import (
	buildconfig "asymmetric-effort/asymmetric-toolkit/buildConfig"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestCopyright(t *testing.T) {
	errors.Assert(buildconfig.Copyright == "(c) 2018 Sam Caldwell.  All Rights Reserved.",
		"expected copyright mismatch")
}
