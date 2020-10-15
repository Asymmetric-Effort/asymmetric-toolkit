package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
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

func TestLoggerHappy(t *testing.T) {
	var log Logger
	errors.Assert(log.level == 0, "expect 0")
	errors.Assert(log.facility == "", "expect empty string")
	errors.Assert(log.writer == nil, "Expect nil writer function pointer")
}
