package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogLevelGetStr_Happy(t *testing.T) {
	func() {
		var o Level = Debug
		errors.Assert(o.String() == "debug", "Expected Debug String")
	}()
	func() {
		var o Level
		errors.Assert(o.String() == "critical", "Expected Critical String")
	}()
}

func TestLogLevelGet_Happy(t *testing.T) {
	var l Level = Critical
	errors.Assert(l.Get() == Critical, "Expect Critical")
	l = Error
	errors.Assert(l.Get() == Error, "Expect Error")
	l = Warning
	errors.Assert(l.Get() == Warning, "Expect Warning")
	l = Info
	errors.Assert(l.Get() == Info, "Expect Info")
	l = Debug
	errors.Assert(l.Get() == Debug, "Expect Debug")
}

func TestLogLevelGet_Sad(t *testing.T) {
	var l Level = BadLevel
	defer func() { recover() }()
	_ = l.Get()
	t.Error("expected error")
}
