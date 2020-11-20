package tags

func (o *String) Add(key string, value string) {
	(*o)[key] = value
}
