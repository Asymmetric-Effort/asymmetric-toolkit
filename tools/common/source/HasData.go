package source

func (o *Source) HasData() bool {
	return len(o.feed) > 0
}
