package sourcetype

func (o *DataSource) IsRandom() bool {
	return *o == Random
}