package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLogWriterStdOut(t *testing.T) {
	var config Configuration
	config.Destination.Set(Stdout)
	config.Level.Set(Debug)

	var log Logger
	log.Setup(&config)

	out := func() string {
		realStdout := os.Stdout
		defer func() { os.Stdout = realStdout }()
		r, fakeStdout, err := os.Pipe()
		exitOnError := func(err error, t *testing.T) {
			if err != nil {
				t.Fatal(err)
			}
		}
		os.Stdout = fakeStdout

		msg := "test"

		log.logWriterStdOut(&msg)

		exitOnError(fakeStdout.Close(), t)
		newOutBytes, err := ioutil.ReadAll(r)
		exitOnError(err, t)
		exitOnError(r.Close(), t)
		return string(newOutBytes)
	}()

	msg := out
	word := msg
	errors.Assert(word == "test", fmt.Sprintf("Expected 'test'. Encountered: '%s'", word))
}
