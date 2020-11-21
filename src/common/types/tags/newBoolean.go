package tags
/*
	tags.Boolean() - create a boolean tag
*/
func NewBoolean() Boolean{
	mutex.Lock()
	defer mutex.Unlock()

	return make(Boolean,1)
}
