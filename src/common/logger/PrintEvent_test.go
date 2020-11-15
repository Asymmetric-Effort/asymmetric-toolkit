package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"encoding/json"
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
	currentTime := time.Now().UnixNano()

	evt := LogEventStruct{
		EventId: EventStd,
		Time:    currentTime,
		Level:   Debug,
		Tags:    []TagId{1, 2, 3, 4},
		Message: msgStr,
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

	var event LogEventStruct
	err := json.Unmarshal([]byte(out), &event)
	if err != nil {
		panic(err)
	}
	//t.Skip("We need to unmarshal the structured log.")
	errors.Assert(out != "{}", fmt.Sprintf("Encountered: '%v'", out))
	errors.Assert(event.EventId == evt.EventId, "expect EventStd")
	errors.Assert(event.Time == evt.Time, fmt.Sprintf("expect time %v == %v", event.Time, evt.Time))
	errors.Assert(event.Level == evt.Level, "expect Debug level")
}
