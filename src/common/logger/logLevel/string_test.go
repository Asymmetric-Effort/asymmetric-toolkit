package logLevel_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogLevelGetStr_Happy(t *testing.T){
	var o LogLevel = Debug
	errors.Assert(o.String()=="DEBUG", "Expected Debug String")
}
