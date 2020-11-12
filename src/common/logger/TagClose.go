package logger

/*
	Logger::TagClose() will emit a log event closing the tagId.  This is an event emission only,
	and it will have no operational effect on the application.  The tag will still be usable at this
	time.  We may find it necessary in the future to panic if tags are used after they are closed.
*/

func (o *Logger) TagClose(tagId TagId) {
	//ToDo: emit a tag event indicating that tagId is closed.
	o.Printf(o.Level, "TagClose(%d)", tagId)
}
