package logLevel_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
	"testing"
)

const (
	BadLevel=-1
)

func TestLogLevelGet_Happy(t *testing.T){
	var l logLevel.LogLevel = logLevel.Critical
	errors.Assert(l.Get() == logLevel.Critical, "Expect Critical")
	l = logLevel.Error
	errors.Assert(l.Get() == logLevel.Error, "Expect Error")
	l = logLevel.Warning
	errors.Assert(l.Get() == logLevel.Warning, "Expect Warning")
	l = logLevel.Info
	errors.Assert(l.Get() == logLevel.Info, "Expect Info")
	l = logLevel.Debug
	errors.Assert(l.Get() == logLevel.Debug, "Expect Debug")
}

func TestLogLevelGet_Sad(t *testing.T){
	var l logLevel.LogLevel = BadLevel
	defer func(){recover()}()
	_ = l.Get()
	t.Error("expected error")
}

