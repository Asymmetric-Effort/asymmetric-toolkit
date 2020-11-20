package tags
/*
	Tag::Delete() - Delete a simple tag and return a boolean value indicating whether or not the key existed.
 */

func (o *Tag)  Delete(key string) (hadRecord bool) {
	if o.Find(key){
		delete(*o,key)
		hadRecord=true
	}
	return
}
