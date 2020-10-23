package fifo

import "asymmetric-effort/asymmetric-toolkit/src/common/errors"

func (o *Fifo) Setup(sz int){
	errors.Assert(sz>0, "expect sz>0")
	o.sz = sz
	o.queue = make(chan string, sz)
}