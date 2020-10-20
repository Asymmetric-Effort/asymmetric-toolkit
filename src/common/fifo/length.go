package fifo

func (o *Fifo) Length() int {
	return len((*o).queue)
}
