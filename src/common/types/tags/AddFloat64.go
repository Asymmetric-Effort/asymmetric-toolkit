package tags

func (o *Float64) Add(key string, value float64) {
	if len(*o) > maxTagLength {
		panic("Too many tags (max:255)")
	}
	(*o)[key] = value
}
