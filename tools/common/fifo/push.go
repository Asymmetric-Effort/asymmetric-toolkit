package fifo

import "time"

func (o *Fifo) Push(s string) {
	for o.Length()+1 > o.sz { //Block if at capacity.
		<-time.After(time.Second*2)
		//ToDo: Emit stat when we block (start, stop times).
	}
	o.queue<-s
}