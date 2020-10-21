package sourcetype

func (o *DataSource) IsSequence() bool {
	return *o == Sequence
}
