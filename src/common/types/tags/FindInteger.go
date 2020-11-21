package tags

func (o *Integer)  Find(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	_, ok := (*o)[key]
	return ok
}
