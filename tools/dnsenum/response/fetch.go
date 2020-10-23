package response

func (o *Collector) Fetch() (data *Response) {
	data=<-o.data
	o.count.requests.Increment()
	if !(*data).connect{
		o.count.connectFail.Increment()
	}
	if (*data).ok {
		o.count.responses.Increment()
	}
	return data
}
