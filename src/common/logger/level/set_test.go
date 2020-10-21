package level_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
	"testing"
)

func TestLogLevelSet_Happy(t *testing.T) {
	var l level.LogLevel
	errors.Assert(l == level.Critical, "Expect Critical (Default)")

	l.Set(level.Error)
	errors.Assert(l.Get() == level.Error, "Expect Error")
	l.Set(level.Warning)
	errors.Assert(l.Get() == level.Warning, "Expect Warningf")
	l.Set(level.Info)
	errors.Assert(l.Get() == level.Info, "Expect Info")
	l.Set(level.Debug)
	errors.Assert(l.Get() == level.Debug, "Expect Debug")
}

func TestLogLevelSet_Sad(t *testing.T) {
	var l level.LogLevel
	defer func() { recover() }()
	l.Set(BadLevel)
	t.Error("expected error")
}
