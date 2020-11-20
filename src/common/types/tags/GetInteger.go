package tags

func (o *Integer) Get(key string) int {
	if o.Find(key) {
		return (*o)[key]
	}
	panic("Key does not exist")
}
