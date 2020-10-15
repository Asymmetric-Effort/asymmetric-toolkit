package counter

func (o *Counter) Increment(p int) bool {
	/*
		Recursively increment the counter from lsb(0) to msb(n)
		and return true if we still have room to increment or
		false if we have reached the counter's maximum value.
	 */
	if (*o.data)[p]<o.maxBase {
		(*o.data)[p]++
		return true //yes, we incremented. You can keep going.
	}else{
		if p<len(*o.data) {
			(*o.data)[p] = 0
			return o.Increment(p+1)
		}else{
			return false //no, we're done.
		}
	}
}
