package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"io/ioutil"
	"os"
	"testing"
)

func exitOnError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func catchStdOut(t *testing.T, runnable func()) string {
	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()
	r, fakeStdout, err := os.Pipe()
	exitOnError(err, t)
	os.Stdout = fakeStdout
	runnable()
	exitOnError(fakeStdout.Close(), t)
	newOutBytes, err := ioutil.ReadAll(r)
	exitOnError(err, t)
	exitOnError(r.Close(), t)
	return string(newOutBytes)
}

func TestLogger(t *testing.T){
	var o Logger
	errors.Assert(o.Level == Critical, "Expect critical(0)")
	errors.Assert(o.Destination == Stdout, "expect stdout(0)")
	errors.Assert(o.Settings == nil, "expect nil")
	errors.Assert(o.Writer == nil, "expect nil")
	errors.Assert(&o.tags != nil, "expect non-nil address")
}