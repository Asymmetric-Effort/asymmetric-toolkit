package logger

/*
	Logger::PrintEvent() consumes a log event and determines if the log event is to be written (evaluating
	log level) and the destination (writer) if configured.
*/
import (
	"encoding/json"
)

func (o *Logger) PrintEvent(event *LogEventStruct) {

	if o.PrintThisLine(event.level) && (o.Writer != nil) {
		event.tags = o.tagMerge(&event.tags)
		msg, err := json.Marshal(*event)
		if err != nil {
			panic(err)
		}
		//fmt.Print(*event)
		out:=string(msg)
		o.Writer(&out)
	}
}
