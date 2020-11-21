package tags

func (o *String) Get(key string) string {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := (*o)[key];ok {
		return (*o)[key]
	}
	panic("Key does not exist")
}
