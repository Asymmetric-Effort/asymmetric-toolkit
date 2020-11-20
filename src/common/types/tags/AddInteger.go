package tags

func (o *Integer) Add(key string, value int) {
	(*o)[key] = value
}
