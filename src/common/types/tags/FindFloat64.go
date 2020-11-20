package tags

func (o *Float64)  Find(key string) bool {
	_, ok := (*o)[key]
	return ok
}
