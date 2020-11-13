package logger

func (o *TagTracker) DeleteGlobal(id TagId) {
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.IsValid(id) {
		delete(o.global, id)
	} else {
		panic("TagTracker::DeleteGlobal() called with invalid TagId")
	}

}
