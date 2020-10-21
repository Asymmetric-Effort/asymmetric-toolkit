package level

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogLevelSet_Happy(t *testing.T){
	var l Level
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
	var l Level
	defer func(){recover()}()
	l.Set(BadLevel)
	t.Error("expected error")
}
