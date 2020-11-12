package logLevel

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

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
