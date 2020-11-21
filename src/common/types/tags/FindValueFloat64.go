package tags

func (o *Float64)  FindValue(key string, value float64) bool {
	mutex.Lock()
	defer mutex.Unlock()

	_, ok := (*o)[key]
	return ok && ((*o)[key] == value)
}
