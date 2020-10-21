package fifo

type Fifo struct {
	sz int
	queue chan string
}
