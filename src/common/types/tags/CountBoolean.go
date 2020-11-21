package tags

func (o *Boolean)Count() int{
	mutex.Lock()
	defer mutex.Unlock()

	return len(*o)
}