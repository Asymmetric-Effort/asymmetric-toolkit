package queues


func (o *Fifo) Pop() (r string) {
	if o.Queue != nil {
		return <-o.Queue
	}else{
		panic("nil queue in Fifo::Pop().  Use .Setup() to initialize.")
	}
}