package reader

func (o *Reader) IsNil() bool {
	return o.file == nil
}