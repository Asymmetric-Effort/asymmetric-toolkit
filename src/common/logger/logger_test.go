package logger_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
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
	var log logger.Logger
	errors.Assert(log.Level == 0, "expect 0")
	errors.Assert(log.Facility == "", "expect empty string")
	errors.Assert(log.Writer == nil, "Expect nil writer function pointer")
}
