package logLevel_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
	"testing"
)

func TestLogLevelGetStr_Happy(t *testing.T){
	var o logLevel.LogLevel = logLevel.Debug
	errors.Assert(o.String()=="DEBUG", "Expected Debug String")
}
