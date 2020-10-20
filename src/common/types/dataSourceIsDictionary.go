package types

func (o *DataSource) IsDictionary() bool {
	return *o == Dictionary
}