package tags

func (o *Float64) Get(key string) float64 {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := (*o)[key];ok {
		return (*o)[key]
	}
	panic("Key does not exist")
}
