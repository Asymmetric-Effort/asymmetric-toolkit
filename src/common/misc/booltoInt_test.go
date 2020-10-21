package misc_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"testing"
)

func TestBoolToInt(t *testing.T) {
	_ = misc.BooltoInt(true)
	_ = misc.BooltoInt(false)
}
