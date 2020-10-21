package level_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
	"testing"
)

const (
	BadLevel=-1
)

func TestLogLevelGet_Happy(t *testing.T){
	var l level.LogLevel = level.Critical
	errors.Assert(l.Get() == level.Critical, "Expect Critical")
	l = level.Error
	errors.Assert(l.Get() == level.Error, "Expect Error")
	l = level.Warning
	errors.Assert(l.Get() == level.Warning, "Expect Warningf")
	l = level.Info
	errors.Assert(l.Get() == level.Info, "Expect Info")
	l = level.Debug
	errors.Assert(l.Get() == level.Debug, "Expect Debug")
}

func TestLogLevelGet_Sad(t *testing.T){
	var l level.LogLevel = BadLevel
	defer func(){recover()}()
	_ = l.Get()
	t.Error("expected error")
}

