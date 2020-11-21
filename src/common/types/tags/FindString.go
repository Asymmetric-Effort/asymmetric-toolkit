package tags

func (o *String)  Find(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	_, ok := (*o)[key]
	return ok
}
