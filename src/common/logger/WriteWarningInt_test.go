package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLogger_WarningInt(t *testing.T) {
	var log Logger
	var config = Configuration{
		Level:       Warning,
		Settings:    nil,
		Destination: Stdout,
	}
	log.Setup(&config)

	tests := []struct {
		Level          Level
		ExpectPrintMsg bool
	}{
		{
			Level:          Debug,
			ExpectPrintMsg: false,
		}, {
			Level:          Info,
			ExpectPrintMsg: false,
		}, {
			Level:          Warning,
			ExpectPrintMsg: true,
		}, {
			Level:          Error,
			ExpectPrintMsg: true,
		}, {
			Level:          Critical,
			ExpectPrintMsg: true,
		},
	}

	for _, test := range tests {
		fmt.Printf("PrintThisLine:\n"+
			"\ttest level     : %v\n"+
			"\tExpectPrintMsg : %v\n"+
			"\tPrintThisLine  : %v\n",
			test.Level.String(),
			test.ExpectPrintMsg,
			log.PrintThisLine(test.Level))

		if test.ExpectPrintMsg {
			if log.PrintThisLine(test.Level) {
				fmt.Println("PrintThisLine() outcome hit: pass")
			} else {
				panic("PrintThisLine() outcome mismatch: pass")
			}
		}
	}
	fmt.Println("")
	fmt.Println("------------")
	fmt.Println("second phase")
	fmt.Println("------------")
	fmt.Println("")
	for _, intMsg := range []int{1, 2, 3} {
		for _, test := range tests {
			var event LogEventStruct

			log.Level = test.Level
			errors.Assert(log.Level == test.Level, fmt.Sprintf("Expect %s", test.Level.String()))

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

				log.WarningInt(EventStd, intMsg)

				exitOnError(fakeStdout.Close(), t)
				newOutBytes, err := ioutil.ReadAll(r)
				exitOnError(err, t)
				exitOnError(r.Close(), t)
				return string(newOutBytes)
			}()

			fmt.Println("\tExpect a print message:", test.ExpectPrintMsg)

			if test.ExpectPrintMsg {
				err := json.Unmarshal([]byte(out), &event)
				if err != nil {
					fmt.Printf("\n"+
						"\tLogLevel: %d\n"+
						"\tError: %v\n"+
						"\tEvent: %v\n",
						test.Level, err, out)
				}
				if event.EventId != EventStd {
					fmt.Printf("Expected '%d'. "+
						"Encountered: '%d'\n", Warning, event.EventId)
				}

				if event.Message == "" {
					event.Message = fmt.Sprintf("%d", intMsg)
				}

				errors.Assert(event.Level <= Warning, "Expected Warning level.")

				fmt.Printf("\nEvent message: '%v'\n", event.Message)

				errors.Assert(event.Message == fmt.Sprintf("%d", intMsg), event.Message)

				fmt.Printf("\tTest with log level '%8s'(%1d): OK\n", test.Level.String(), test.Level)
			}
		}
	}
}
