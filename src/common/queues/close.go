package queues

func (o *Fifo) Close(){
	if o.Queue == nil {
		panic("Cannot close nil queue")
	}
	close(o.Queue)
}