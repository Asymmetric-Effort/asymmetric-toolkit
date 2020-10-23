package types

func (o *DataSource) IsNotSet() bool {
	return *o == NotSet
}
