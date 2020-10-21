package errors_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestWarn(t *testing.T){
	errors.Warn("pass") //cause assertion error."
}
