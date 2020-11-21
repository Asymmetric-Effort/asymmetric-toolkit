package tags

func (o *Integer) Get(key string) int {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := (*o)[key];ok {
		return (*o)[key]
	}
	panic("Key does not exist")
}
