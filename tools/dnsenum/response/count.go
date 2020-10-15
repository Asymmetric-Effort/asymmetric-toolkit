package response

type Count struct {
	requests    Counter //Number requests sent
	connectFail Counter //Number of connections failed
	responses   Counter //Number of responses rec'd
}

type Counter int

func (o *Counter) Increment(){
	*o++
}

func (o *Counter) Get() int {
	return int(*o)
}