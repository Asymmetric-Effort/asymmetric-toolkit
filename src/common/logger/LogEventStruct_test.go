package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestLogEventStruct(t *testing.T) {
	for level := Level(Critical); level <= Level(Debug); level++ {
		for id := EventId(0); id <= EventId(3); id++ {

			currTime := time.Now().UnixNano()

			var src LogEventStruct = LogEventStruct{
				EventId: id,
				Time:    currTime,
				Level:   level,
				Tags:    []TagId{},
				Message: "",
			}
			var blob []byte
			var err error

			fmt.Printf("\tlevel:%v, id:%v, time:%v\n", level, id, currTime)

			func() {
				errors.Assert(src.EventId == id, fmt.Sprintf("expect eventId %v == %v", src.EventId, id))
				errors.Assert(src.Time == currTime, fmt.Sprintf("expect time (pre) %v", src.Time))
				errors.Assert(src.Level == level,
					fmt.Sprintf("expect level %v (%v)", src.Level.String(), src.Level))
				errors.Assert(src.Message == "", fmt.Sprintf("expect message %v", src.Message))
				blob, err = json.Marshal(src)
				if err != nil {
					panic(err)
				}
			}()

			func() { // Unmarshall JSON blob into struct
				var dest LogEventStruct
				err = json.Unmarshal(blob, &dest)
				errors.Assert(dest.EventId == id, "expect EventStd")
				errors.Assert(dest.Level == level, "expect critical (0)")
				errors.Assert(dest.Message == "", "expect message ''")
				errors.Assert(src.Time == currTime,
					fmt.Sprintf("expect %v in pre-test %v", src.Time, currTime))

				errors.Assert(dest.Time == currTime,
					fmt.Sprintf("expect time (post) %v | %v", dest.Time, currTime))

				if err != nil {
					panic(err)
				}
			}()
		}
	}

}
