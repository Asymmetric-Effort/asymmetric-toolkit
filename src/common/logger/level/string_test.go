package level

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogLevelGetStr_Happy(t *testing.T){
	var o Level = Debug
	errors.Assert(o.String()=="DEBUG", "Expected Debug String")
}
