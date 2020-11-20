package tags
/*
	Tag::Add() - Add a simple tag.
 */

func (o *Tag) Add(key string) {
	(*o)[key] = struct{}{}
}
