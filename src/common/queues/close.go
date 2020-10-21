package queues

func (o *Fifo) Close(){
	if o.queue == nil {
		panic("Cannot close nil queue")
	}
	close(o.queue)
}