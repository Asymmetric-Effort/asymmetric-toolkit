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
		event.tags = o.tagMerge(event.tags)
		msg, err := json.Marshal(*event)
		if err != nil {
			panic(err)
		}
		o.Writer(&msg)
	}
}

func (o *Logger) tagMerge(tags *[]TagId) *[]TagId {
	var mergedTags = make(TagTable, 1)
	for tag, _ := range o.tags.global {
		mergedTags[tag] = struct{}{}
	}
	if tags != nil {
		for _, tag := range *tags {
			mergedTags[tag] = struct{}{}
		}
	}
	var result []TagId
	for tag,_:=range mergedTags{
		result=append(result,tag)
	}
	return &result
}
