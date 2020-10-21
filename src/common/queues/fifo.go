package queues

type Fifo struct {
	sz    int
	Queue chan string
}
