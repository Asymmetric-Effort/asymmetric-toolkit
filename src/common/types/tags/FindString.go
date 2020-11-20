package tags

func (o *String)  Find(key string) bool {
	_, ok := (*o)[key]
	return ok
}
