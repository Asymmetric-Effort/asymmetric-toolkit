package response

func (o *Collector) Setup(buffSz int) {
	o.data = make(chan *Response, buffSz)
}