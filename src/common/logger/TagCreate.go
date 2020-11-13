package logger

/*
	Logger::TagCreate() creates a new TagId and emits a descriptive event.
*/

import (
	"fmt"
)

func (o *Logger) TagCreate(tagName string) (tagId TagId) {
	tagId = o.tags.Create(&tagName)
	tagData := fmt.Sprintf("%s:%d", tagName, tagId)
	o.Printf(AnyLevel, EventTagCreate, nil, &tagData)
	return tagId
}
