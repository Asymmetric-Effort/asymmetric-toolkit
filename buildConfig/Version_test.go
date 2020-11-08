package buildconfig_test

import (
	buildconfig "asymmetric-effort/asymmetric-toolkit/buildConfig"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestVersion(t *testing.T) {
	errors.Assert(buildconfig.Version == "0.0.1", "expected version mismatch")
}
