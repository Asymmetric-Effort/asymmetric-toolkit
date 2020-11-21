package tags

func (o *Integer) Add(key string, value int) {
	mutex.Lock()
	defer mutex.Unlock()

	if len(*o) > maxTagLength {
		panic("Too many tags (max:255)")
	}
	(*o)[key] = value
}
