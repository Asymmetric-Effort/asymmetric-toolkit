package response

import "time"

type Response struct {
	startTime time.Time
	//stopTime  time.Time
	connect   bool // Server connected ok
	ok        bool // Server responded ok
	//data      []byte
}

type New struct {}

func (o New) NewResponse() *Response{
	var r Response
		r.startTime=time.Now()
	return &r
}