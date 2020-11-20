package tags

func (o *String) Add(key string, value string) {
	if len(*o) > maxTagLength {
		panic("Too many tags (max:255)")
	}
	(*o)[key] = value
}
