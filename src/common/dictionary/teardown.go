package dictionary



func (o *Dictionary) Teardown() {
	if o.fileHandle != nil {
		err := o.fileHandle.Close()
		if err != nil {
			panic(err)
		}
	}
}