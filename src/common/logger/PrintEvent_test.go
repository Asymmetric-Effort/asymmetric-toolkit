package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestLogger_PrintEvent(t *testing.T) {
	var config Configuration
	config.Destination.Set(Stdout)
	config.Level.Set(Debug)

	var log Logger
	log.Setup(&config)
	msgStr := "testMessage"

	evt := LogEventStruct{
		eventId: EventStd,
		time:    time.Now(),
		level:   Debug,
		tags:    []TagId{1, 2, 3, 4},
		message: msgStr,
	}

	exitOnError := func(err error, t *testing.T) {
		if err != nil {
			t.Fatal(err)
		}
	}

	out := func() string {
		realStdout := os.Stdout
		defer func() { os.Stdout = realStdout }()
		r, fakeStdout, err := os.Pipe()
		exitOnError(err, t)
		os.Stdout = fakeStdout

		log.PrintEvent(&evt)

		exitOnError(fakeStdout.Close(), t)
		newOutBytes, err := ioutil.ReadAll(r)
		exitOnError(err, t)
		exitOnError(r.Close(), t)
		return string(newOutBytes)
	}()
	fmt.Println("out:", out)
	//errors.Assert(out != "{}", fmt.Sprintf("Encountered: '%v'", out))
}
