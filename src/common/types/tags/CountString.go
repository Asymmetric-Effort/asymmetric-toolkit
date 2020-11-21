package tags

func (o *String)Count() int{
	mutex.Lock()
	defer mutex.Unlock()

	return len(*o)
}