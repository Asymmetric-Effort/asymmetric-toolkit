package level_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
	"testing"
)

func TestLogLevelGetStr_Happy(t *testing.T){
	var o level.LogLevel = level.Debug
	errors.Assert(o.String()=="DEBUG", "Expected Debug String")
}
