package logger

/*
	Logger::Printf() is the log print formatter method which formats the log event content as a structured JSON
	object before passing the result to Logger::Print() for transmission to the configured destination.
*/
import (
	"fmt"
	"time"
)

const formatString = `{"t":%d,"s":%d,"f":%d,"e":%d,"v":"%d","m":"%s","tags":[%s]}`

/*
	t    - time
	s    - severity (log level)
	f    - facility (log facility)
	e    - eventId  (log eventId)
	v    - value (numeric log value) (Optional.  Default:0)
	m    - message string.
	tags - arbitrary tagIds.
*/

func (o *Logger) Printf(lvl Level, eventId EventId, tags *[]TagId, msg *string, value ...int) {

	o.Print(lvl, fmt.Sprintf(formatString,
		time.Now().UnixNano(),
		lvl.Get(),
		o.Facility.Get(),
		eventId,
		func() int {
			//Evaluate any value we might expect.
			switch len(value) {
			case 0:
				return 0
			case 1:
				return value[0]
			default:
				panic("only one value may be passed to Logger::Printf")
			}
		}(),
		func() string {
			// Validate the arbitrary message string.
			if msg == nil {
				return ""
			} else {
				if len([]byte(*msg)) > maxLogMessageStringLength {
					panic(fmt.Sprintf("Log messages should be < %d bytes",maxLogMessageStringLength))
				}
				return *msg
			}
		}(),
		func() (s string) {
			//Stringify tags.
			if tags == nil {
				return ""
			}
			format := `%s,"%d"`
			for _, tag := range *tags {
				s = fmt.Sprintf(format, s, tag)
			}
			return s
		}()))
}
