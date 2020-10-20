package types

func (o *DataSource) IsRandom() bool {
	return *o == Random
}