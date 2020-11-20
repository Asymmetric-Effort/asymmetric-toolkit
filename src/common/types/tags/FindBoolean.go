package tags

func (o *Boolean)  Find(key string) bool {
	_, ok := (*o)[key]
	return ok
}
