package tags

func (o *String)  FindValue(key string, value string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	_, ok := (*o)[key]
	return ok && ((*o)[key] == value)
}
