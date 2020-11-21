package tags

func (o *Float64)  Delete(key string) (hadRecord bool) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := (*o)[key];ok {
		delete(*o,key)
		hadRecord=true
	}
	return
}
