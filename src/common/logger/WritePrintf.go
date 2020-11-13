package logger

/*
	Logger::Printf() is the log print formatter method which formats the log event content as a structured JSON
	object before passing the result to Logger::Print() for transmission to the configured destination.
*/
import (
	"fmt"
	"time"
)

const formatString = `{"t":%d,"s":%d,"e":%d,"m":"%s","tags":[%s]}`

/*
	t    - time
	s    - severity (log level)
	e    - eventId  (log eventId)
	v    - value (numeric log value) (Optional.  Default:0)
	m    - message string.
	tags - arbitrary tagIds.
*/

func (o *Logger) Printf(lvl Level, eventId EventId, tags *[]TagId, msg ...*string) {
	o.Print(lvl, fmt.Sprintf(formatString,
		time.Now().UnixNano(),
		lvl.Get(),
		eventId,
		func() string {
			// Validate the arbitrary message string.
			if msg == nil {
				return ""
			} else {
				if len([]byte(*msg[0])) > maxLogMessageStringLength {
					panic(fmt.Sprintf("Log messages should be < %d bytes", maxLogMessageStringLength))
				}
				return *msg[0]
			}
		}(),
		func() (s string) {
			//Stringify tags.
			if tags == nil {
				return ""
			}
			const format = `%s,"%d"`
			var localTags = make(map[TagId]bool, 1)
			for tag, _ := range o.tags.global { //append global tags to our local
				localTags[tag] = true
			}
			for tag, _ := range localTags {
				s = fmt.Sprintf(format, s, tag)
			}
			return s
		}()))
}
