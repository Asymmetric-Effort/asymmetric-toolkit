package logLevel_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

const (
	BadLevel=-1
)

func TestLogLevelGet_Happy(t *testing.T){
	var l LogLevel = Critical
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

func TestLogLevelGet_Sad(t *testing.T){
	var l LogLevel = BadLevel
	defer func(){recover()}()
	_ = l.Get()
	t.Error("expected error")
}

