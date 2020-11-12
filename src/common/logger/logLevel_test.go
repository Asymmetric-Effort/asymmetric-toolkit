package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"strings"
	"testing"
)

const (
	BadLevel=-1
)

func TestLogLevel_Happy(t *testing.T) {
	errors.Assert(Critical == 0, "Expect 0")
	errors.Assert(Error == 1, "Expect 1")
	errors.Assert(Warning == 2, "Expect 2")
	errors.Assert(Info == 3, "Expect 3")
	errors.Assert(Debug == 4, "Expect 4")
	errors.Assert(LevelStrings == "critical,error,warning,info,debug", "LevelStrings has changed or has error")
	errors.Assert(len(strings.Split(LevelStrings, ",")) == 5, "Expected 5 level strings.  Found more or less")
}

func TestLogLevelSet_Happy(t *testing.T){
	var l LogLevel
	errors.Assert(l == Critical, "Expect Critical (Default)")

	l.Set(Error)
	errors.Assert(l.Get() == Error, "Expect Error")
	l.Set(Warning)
	errors.Assert(l.Get() == Warning, "Expect Warning")
	l.Set(Info)
	errors.Assert(l.Get() == Info, "Expect Info")
	l.Set(Debug)
	errors.Assert(l.Get() == Debug, "Expect Debug")
}

func TestLogLevelSet_Sad(t *testing.T){
	var l LogLevel
	defer func(){recover()}()
	l.Set(BadLevel)
	t.Error("expected error")
}

func TestLogLevelGetStr_Happy(t *testing.T){
	func(){
		var o LogLevel = Debug
		errors.Assert(o.String()=="debug", "Expected Debug String")
	}()
	func(){
		var o LogLevel
		errors.Assert(o.String()=="critical", "Expected Critical String")
	}()
}

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

