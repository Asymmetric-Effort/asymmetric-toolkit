package queues

type Fifo struct {
	sz int
	queue chan string
}
