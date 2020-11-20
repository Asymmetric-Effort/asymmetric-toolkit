package tags

func (o *String) Get(key string) string {
	if o.Find(key) {
		return (*o)[key]
	}
	panic("Key does not exist")
}
