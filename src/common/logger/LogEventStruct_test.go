package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"encoding/json"
	"testing"
	"time"
)

func TestLogEventStruct(t *testing.T) {
	var o LogEventStruct
	var blob []byte
	var err error

	for level := Level(Critical); level <= Level(Debug); level++ {
		for id := EventId(0); id <= EventId(3); id++ {

			o.eventId = id
			o.time = time.Time{}
			o.level = level

			func() { // Verify that all expected fields are in initial state.
				errors.Assert(o.eventId == id, "expect 0")
				errors.Assert(o.time == time.Time{}, "expect 0")
				errors.Assert(o.level == level, "expect critical (0)")
				errors.Assert(o.tags == nil, "expect nil")
				errors.Assert(o.message == "", "expect nil")
			}()

			func() { // Marshall struct to JSON
				blob, err = json.Marshal(o)
				if err != nil {
					panic(err)
				}
				errors.Assert(string(blob) != "", "expect non-empty string.")
			}()

			func() { // Unmarshall JSON blob into struct
				var event LogEventStruct
				err = json.Unmarshal(blob, &event)
				if err != nil {
					panic(err)
				}
			}()
		}
	}

}
