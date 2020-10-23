package dictionary

func (o *Dictionary) DefCount() int{
	return int(o.content.body.defCount)
}
