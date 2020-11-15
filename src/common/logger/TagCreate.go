package logger

/*
	Logger::TagCreate() creates a new TagId and emits a descriptive event.
*/

import (
	"time"
)

func (o *Logger) TagCreate(tagName string) (tagId TagId) {
	tagId = o.tags.Create(tagName)
	o.PrintEvent(&LogEventStruct{
		EventId: EventTagCreate,
		Time:    time.Now().UnixNano(),
		Level:   Any,
		Tags:    []TagId{tagId},
		Message: tagName,
	})
	return tagId
}
