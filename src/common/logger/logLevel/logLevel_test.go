package logLevel_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
	"strings"
	"testing"
)

func TestLogLevel_Happy(t *testing.T) {
	errors.Assert(logLevel.Critical == 0, "Expect 0")
	errors.Assert(logLevel.Error == 1, "Expect 1")
	errors.Assert(logLevel.Warning == 2, "Expect 2")
	errors.Assert(logLevel.Info== 3, "Expect 3")
	errors.Assert(logLevel.Debug == 4, "Expect 4")
	errors.Assert(logLevel.LevelStrings == "CRIT,ERROR,WARN,INFO,DEBUG", "LevelStrings has changed or has error")
	errors.Assert(len(strings.Split(logLevel.LevelStrings, ",")) == 5, "Expected 5 level strings.  Found more or less")
}
