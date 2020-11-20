package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestTag_Constants(t *testing.T) {
	errors.Assert(maxTagLength == 255, "maxTagLength mistmatch.")
}
