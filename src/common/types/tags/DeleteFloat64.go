package tags

func (o *Float64)  Delete(key string) (hadRecord bool) {
	if o.Find(key){
		delete(*o,key)
		hadRecord=true
	}
	return
}
