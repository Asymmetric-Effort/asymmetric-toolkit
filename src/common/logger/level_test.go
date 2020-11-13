package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"math"
	"strings"
	"testing"
)

const (
	BadLevel = -1
)

func TestLogLevel(t *testing.T) {
	func() {
		errors.Assert(Critical == 0, "Expect 0")
		errors.Assert(Error == 1, "Expect 1")
		errors.Assert(Warning == 2, "Expect 2")
		errors.Assert(Info == 3, "Expect 3")
		errors.Assert(Debug == 4, "Expect 4")
		errors.Assert(LevelStrings == "critical,error,warning,info,debug", "LevelStrings has changed or has error")
		errors.Assert(len(strings.Split(LevelStrings, ",")) == 5, "Expected 5 level strings.  Found more or less")
	}()
	func() {
		defer func() { recover() }()
		var o Level = -1
		_ = o.Get()
		panic("Expected error.")
	}()
	func(){
		defer func() { recover() }()
		var o Level = 10
		_=o.Get()
		panic("Expected error.")
	}()
	func(){
		defer func() { recover() }()
		var o Level = math.MaxInt32
		_=o.Get()
		panic("Expected error.")
	}()
}
