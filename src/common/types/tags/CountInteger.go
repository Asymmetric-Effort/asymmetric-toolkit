package tags

func (o *Integer)Count() int{
	mutex.Lock()
	defer mutex.Unlock()

	return len(*o)
}