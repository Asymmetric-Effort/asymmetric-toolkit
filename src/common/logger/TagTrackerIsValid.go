package logger
/*
	TagTracker::IsValid() returns a boolean indicating whether or not a given TagId is known to the
	tracker object.
 */
func (o *TagTracker) IsValid(id TagId) bool {
	if o.tagIds == nil {
		panic("TagTracker must be initialized before IsValid() is called.")
	}
	_, ok := (*o).tagIds[id]
	return ok
}
