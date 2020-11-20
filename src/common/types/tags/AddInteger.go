package tags

func (o *Integer) Add(key string, value int) {
	if len(*o) > maxTagLength {
		panic("Too many tags (max:255)")
	}
	(*o)[key] = value
}
