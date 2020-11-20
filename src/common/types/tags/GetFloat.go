package tags

func (o *Float64) Get(key string) float64 {
	if o.Find(key) {
		return (*o)[key]
	}
	panic("Key does not exist")
}
