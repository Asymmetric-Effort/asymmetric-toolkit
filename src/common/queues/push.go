package queues

import "time"

func (o *Fifo) Push(s string) {
	if o.Queue != nil {
		for o.Length()+1 > o.sz { //Block if at capacity.
			<-time.After(QueueWriteBlockDelay)
			//ToDo: Emit stat when we block (start, stop times).
		}

		o.Queue <- s
	}else{
		panic("nil queue in Fifo::Push().  Use .Setup() to initialize.")
	}
}