package tags
/*
	Tag::Delete() - Delete a simple tag and return a boolean value indicating whether or not the key existed.
 */

func (o *Tag)  Delete(key string) (hadRecord bool) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := (*o)[key];ok {
		delete(*o,key)
		hadRecord=true
	}
	return
}
