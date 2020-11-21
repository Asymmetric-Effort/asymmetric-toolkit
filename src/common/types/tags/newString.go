package tags

func NewString() String{
	mutex.Lock()
	defer mutex.Unlock()

	return make(String,1)
}
