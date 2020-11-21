package tags

func (o *Integer)  FindValue(key string, value int) bool {
	mutex.Lock()
	defer mutex.Unlock()

	_, ok := (*o)[key]
	return ok && ((*o)[key] == value)
}
