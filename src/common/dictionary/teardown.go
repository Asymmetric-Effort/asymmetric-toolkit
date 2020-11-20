package dictionary



func (o *Descriptor) Teardown() {
	if o.fileHandle != nil {
		err := o.fileHandle.Close()
		if err != nil {
			panic(err)
		}
	}
}