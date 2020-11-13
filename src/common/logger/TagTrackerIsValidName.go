package logger

/*
	TagTracker::IsValidName() returns a boolean indicating whether or not a given tag name is known to the
	tracker object.
 */

func (o *TagTracker) IsValidName(name string) bool {
	_, ok := (*o).tagNames[name]
	return ok
}
