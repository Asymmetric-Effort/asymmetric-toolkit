package tags

func (o *Boolean)Get(key string) bool {
	if o.Find(key){
		return (*o)[key]
	}
	panic("Key does not exist")
}