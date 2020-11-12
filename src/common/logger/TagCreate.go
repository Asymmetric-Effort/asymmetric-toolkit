package logger

/*
	Logger::TagCreate() creates a new TagId
*/

func (o *Logger) TagCreate(tagName string) (tagId TagId) {
	o.nextTag++
	//ToDo: emit a tag event relating the associated TagId and TagName.
	return o.nextTag
}
