package logLevel

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestLogLevelGetStr_Happy(t *testing.T){
	var o LogLevel = Debug
	errors.Assert(o.String()=="DEBUG", "Expected Debug String")
}
