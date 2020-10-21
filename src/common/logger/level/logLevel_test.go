package level_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
	"strings"
	"testing"
)

func TestLogLevel_Happy(t *testing.T) {
	errors.Assert(level.Critical == 0, "Expect 0")
	errors.Assert(level.Error == 1, "Expect 1")
	errors.Assert(level.Warning == 2, "Expect 2")
	errors.Assert(level.Info== 3, "Expect 3")
	errors.Assert(level.Debug == 4, "Expect 4")
	errors.Assert(level.LevelStrings == "CRIT,ERROR,WARN,INFO,DEBUG", "LevelStrings has changed or has error")
	errors.Assert(len(strings.Split(level.LevelStrings, ",")) == 5, "Expected 5 level strings.  Found more or less")
}
