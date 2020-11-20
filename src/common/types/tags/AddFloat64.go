package tags

func (o *Float64) Add(key string, value float64) {
	(*o)[key] = value
}
