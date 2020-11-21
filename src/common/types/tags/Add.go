package tags
/*
	Tag::Add() - Add a simple tag.
 */

func (o *Tag) Add(key string) {
	mutex.Lock()
	defer mutex.Unlock()

	if len(*o) > maxTagLength {
		panic("Too many tags (max:255)")
	}
	(*o)[key] = struct{}{}
}
