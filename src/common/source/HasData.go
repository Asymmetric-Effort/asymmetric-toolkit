package source

func (o *Source) HasData() bool {
	return o.feed.Length() > 0
}
