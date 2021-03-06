package logger

func (o *TagTracker) Close(id TagId) {
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.tagIds == nil {
		panic("TagTracker must be initialized before Close() is called.")
	}
	if o.IsValid(id) {
		if _, ok := o.global[id]; ok {
			delete(o.global, id)
		} // If not found, move on.  Not all will be global.
		delete(o.tagNames, o.tagIds[id])
		delete(o.tagIds, id)
	}else{
		panic("TagTracker::Close() attempted to close a non-existent tag.")
	}
}
