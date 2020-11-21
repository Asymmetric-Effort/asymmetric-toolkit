package tags
/*
	tags.NewTag() - create an empty tag
 */
func NewTag() Tag{
	mutex.Lock()
	defer mutex.Unlock()

	return make(Tag,1)
}