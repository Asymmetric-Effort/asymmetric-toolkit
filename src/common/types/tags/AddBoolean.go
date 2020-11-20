package tags

func (o *Boolean) Add(key string, value bool) {
	(*o)[key] = value
}
