package tags

func (o *String)  FindValue(key string, value string) bool {
	_, ok := (*o)[key]
	return ok && ((*o)[key] == value)
}
