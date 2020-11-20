package tags

func (o *Integer)  FindValue(key string, value int) bool {
	_, ok := (*o)[key]
	return ok && ((*o)[key] == value)
}
