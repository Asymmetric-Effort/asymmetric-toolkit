package fifo


func (o *Fifo) Pop() (r string) {
	return <-(*o).queue
}