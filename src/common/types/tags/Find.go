package tags
/*
	Tag::Find() - Return true/false if a Tag key exists or not.
 */
func (o *Tag)  Find(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	_, ok := (*o)[key]
	return ok
}
