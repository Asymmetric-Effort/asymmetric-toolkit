package tags

func (o *Float64)Count() int{
	mutex.Lock()
	defer mutex.Unlock()

	return len(*o)
}