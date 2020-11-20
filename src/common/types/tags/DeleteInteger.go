package tags

func (o *Integer)  Delete(key string) (hadRecord bool) {
	if o.Find(key){
		delete(*o,key)
		hadRecord=true
	}
	return
}
