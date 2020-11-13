package logger

func (o *TagTracker) AddGlobal(id TagId){
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.IsValid(id){
		o.global[id]=true
	}else{
		panic("TagTracker::AddGlobal() called with invalid TagId")
	}
}