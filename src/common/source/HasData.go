package source

func (o *Source) HasData() bool {
	return o.Feed.Length() > 0
}
