package tags

func (o *Integer)  Find(key string) bool {
	_, ok := (*o)[key]
	return ok
}
