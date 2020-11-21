package tags

func NewInteger() Integer{
	mutex.Lock()
	defer mutex.Unlock()

	return make(Integer,1)
}
