package tags

func (o *Boolean)Get(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := (*o)[key];ok {
		return (*o)[key]
	}
	panic("Key does not exist")
}