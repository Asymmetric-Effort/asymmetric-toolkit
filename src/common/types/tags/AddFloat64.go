package tags

func (o *Float64) Add(key string, value float64) {
	mutex.Lock()
	defer mutex.Unlock()

	if len(*o) > maxTagLength {
		panic("Too many tags (max:255)")
	}
	(*o)[key] = value
}
