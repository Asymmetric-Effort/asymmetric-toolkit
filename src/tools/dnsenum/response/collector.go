package response

type Collector struct {
	count Count
	data  chan *Response
}