package logLevel_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
	"testing"
)

func TestLogLevelSet_Happy(t *testing.T) {
	var l logLevel.LogLevel
	errors.Assert(l == logLevel.Critical, "Expect Critical (Default)")

	l.Set(logLevel.Error)
	errors.Assert(l.Get() == logLevel.Error, "Expect Error")
	l.Set(logLevel.Warning)
	errors.Assert(l.Get() == logLevel.Warning, "Expect Warning")
	l.Set(logLevel.Info)
	errors.Assert(l.Get() == logLevel.Info, "Expect Info")
	l.Set(logLevel.Debug)
	errors.Assert(l.Get() == logLevel.Debug, "Expect Debug")
}

func TestLogLevelSet_Sad(t *testing.T) {
	var l logLevel.LogLevel
	defer func() { recover() }()
	l.Set(BadLevel)
	t.Error("expected error")
}
