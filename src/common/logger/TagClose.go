package logger

/*
	Logger::TagClose() will emit a log event closing the tagId.  This is an event emission only,
	and it will have no operational effect on the application.  The tag will still be usable at this
	time.  We may find it necessary in the future to panic if tags are used after they are closed.
*/

import (
	"fmt"
)

func (o *Logger) TagClose(tagId TagId) {
	o.tags.Close(tagId)
	tagData := fmt.Sprintf("%d", tagId)
	o.Printf(AnyLevel, EventTagClose, nil, &tagData)
}
