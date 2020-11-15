package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLogger_Critical(t *testing.T) {
	var log Logger

	tests := []struct {
		level Level
		pass  bool
	}{
		{
			level: Critical,
			pass:  true,
		}, {
			level: Error,
			pass:  false,
		}, {
			level: Warning,
			pass:  false,
		}, {
			level: Info,
			pass:  false,
		}, {
			level: Debug,
			pass:  false,
		},
	}

	for _, test := range tests {
		var config = Configuration{
			Level:       test.level,
			Settings:    nil,
			Destination: Stdout,
		}
		log.Setup(&config)
		errors.Assert(log.Level == test.level, fmt.Sprintf("Expect %s", test.level.String()))

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

			log.Critical(EventStd)

			exitOnError(fakeStdout.Close(), t)
			newOutBytes, err := ioutil.ReadAll(r)
			exitOnError(err, t)
			exitOnError(r.Close(), t)
			return string(newOutBytes)
		}()
		if test.pass {
			var event LogEventStruct
			err := json.Unmarshal([]byte(out), &event)
			if err != nil {
				panic(fmt.Sprintf("Error:%v  (%v)",err, event))
			}
			errors.Assert(event.eventId == EventStd, fmt.Sprintf("Expected '0'. Encountered: '%d'", event.eventId))
		} else {
			errors.Assert(out == "", fmt.Sprintf("Expected no log written for level %d.", test.level))
		}
		fmt.Printf("\tTest with log level '%8s'(%1d): OK\n", test.level.String(), test.level)
	}
}
