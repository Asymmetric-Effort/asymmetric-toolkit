package source

func (o *Source) HasData() bool {
	return o.queue.Length() > 0
}
