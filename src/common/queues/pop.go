package queues


func (o *Fifo) Pop() (r string) {
	if o.queue != nil {
		return <-o.queue
	}else{
		panic("nil queue in Fifo::Pop().  Use .Setup() to initialize.")
	}
}