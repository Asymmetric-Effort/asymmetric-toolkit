package response

func (o *Collector) Collect(response *Response) {
	o.data<-response
}