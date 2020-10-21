package counter

func (o *Counter) Increment(p int) bool {
	/*
		Recursively increment the counter from lsb(0) to msb(n)
		and return true if we still have room to increment or
		false if we have reached the counter's maximum value.
	*/
	if p >= len(*o.Data) {
		return false
	}
	if (*o.Data)[p] >= o.MaxBase {
		(*o.Data)[p] = 0
		return o.Increment(p + 1) //Carry the one recursively
	}
	(*o.Data)[p]++
	return true //yes, we incremented. You can keep going.
}
