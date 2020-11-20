package tags
/*
	tags.Boolean::FindValue(key,value) - Find a matching tag by key and value
 */
func (o *Boolean)  FindValue(key string, value bool) bool {
	_, ok := (*o)[key]
	return ok && ((*o)[key] == value)
}
