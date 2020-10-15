package types

func (o *DataSource) IsSequence() bool {
	return *o == Sequence
}
