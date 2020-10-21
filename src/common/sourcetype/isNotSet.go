package sourcetype

func (o *DataSource) IsNotSet() bool {
	return *o == NotSet
}
