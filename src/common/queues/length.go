package queues

func (o *Fifo) Length() int {
	return len((*o).Queue)
}
