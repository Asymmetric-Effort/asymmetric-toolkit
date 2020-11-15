package logger

func (o *TagTracker) AddGlobal(id TagId){
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.global == nil {
		o.global=make(TagTable,1)
	}
	if o.IsValid(id){
		o.global[id]=struct{}{}
	}else{
		panic("TagTracker::AddGlobal() called with invalid TagId")
	}
}