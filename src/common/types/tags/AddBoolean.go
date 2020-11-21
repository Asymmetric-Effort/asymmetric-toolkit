package tags

func (o *Boolean) Add(key string, value bool) {
	mutex.Lock()
	defer mutex.Unlock()

	if len(*o) > maxTagLength {
		panic("Too many tags (max:255)")
	}
	(*o)[key] = value
}
