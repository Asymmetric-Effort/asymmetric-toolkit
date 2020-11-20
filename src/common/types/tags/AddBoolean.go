package tags

func (o *Boolean) Add(key string, value bool) {
	if len(*o) > maxTagLength {
		panic("Too many tags (max:255)")
	}
	(*o)[key] = value
}
