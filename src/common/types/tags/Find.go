package tags
/*
	Tag::Find() - Return true/false if a Tag key exists or not.
 */
func (o *Tag)  Find(key string) bool {
	_, ok := (*o)[key]
	return ok
}
