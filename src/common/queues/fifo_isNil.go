package queues

func (o *Fifo) IsNil() bool{
	return o.queue == nil
}