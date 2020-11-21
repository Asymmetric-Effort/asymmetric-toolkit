package tags

func (o *Float64)  Find(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	_, ok := (*o)[key]
	return ok
}
