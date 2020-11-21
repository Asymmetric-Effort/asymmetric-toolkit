package tags

func (o *Boolean)  Find(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	_, ok := (*o)[key]
	return ok
}
