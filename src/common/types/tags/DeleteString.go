package tags

func (o *String)  Delete(key string) (hadRecord bool) {
	if o.Find(key){
		delete(*o,key)
		hadRecord=true
	}
	return
}
