package logger

/*
	Logger::TagCreate() creates a new TagId and emits a descriptive event.
*/

import (
	"time"
)

func (o *Logger) TagCreate(tagName string) (tagId TagId) {
	tagId = o.tags.Create(&tagName)
	o.PrintEvent(&LogEventStruct{
		eventId: EventTagCreate,
		time:    time.Now(),
		level:   Any,
		tags:    []TagId{tagId},
		message: tagName,
	})
	return tagId
}
